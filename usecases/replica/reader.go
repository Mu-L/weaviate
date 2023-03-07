//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package replica

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/storobj"
	"github.com/weaviate/weaviate/usecases/objects"
)

type reader struct {
	class  string
	client finderClient // needed to commit and abort operation
	repairer
	log logrus.FieldLogger
}

type result[T any] struct {
	data T
	err  error
}

type tuple[T any] struct {
	sender string
	UTime  int64
	o      T
	ack    int
	err    error
}

type (
	objTuple  tuple[objects.Replica]
	boolTuple tuple[RepairResponse]
)

type vote struct {
	batchReply
	Count []int
	Err   error
}

func (f *reader) readOne(ctx context.Context,
	shard string,
	id strfmt.UUID,
	ch <-chan simpleResult[findOneReply],
	st rState,
) <-chan result[*storobj.Object] {
	// counters tracks the number of votes for each participant
	resultCh := make(chan result[*storobj.Object], 1)
	go func() {
		defer close(resultCh)
		var (
			votes      = make([]objTuple, 0, len(st.Hosts))
			maxCount   = 0
			contentIdx = -1
		)

		for r := range ch { // len(ch) == st.Level
			resp := r.Response
			if r.Err != nil { // a least one node is not responding
				f.log.WithField("op", "get").WithField("replica", resp.sender).
					WithField("class", f.class).WithField("shard", shard).WithField("uuid", id).Error(r.Err)
				resultCh <- result[*storobj.Object]{nil, errRead}
				return
			}
			if !resp.DigestRead {
				contentIdx = len(votes)
			}
			votes = append(votes, objTuple{resp.sender, resp.UpdateTime, resp.Data, 0, nil})
			for i := range votes { // count number of votes
				if votes[i].UTime == resp.UpdateTime {
					votes[i].ack++
				}
				if maxCount < votes[i].ack {
					maxCount = votes[i].ack
				}
				if maxCount >= st.Level && contentIdx >= 0 {
					resultCh <- result[*storobj.Object]{votes[contentIdx].o.Object, nil}
					return
				}
			}
		}

		obj, err := f.repairOne(ctx, shard, id, votes, st, contentIdx)
		if err == nil {
			resultCh <- result[*storobj.Object]{obj, nil}
			return
		}

		resultCh <- result[*storobj.Object]{nil, errRepair}
		var sb strings.Builder
		for i, c := range votes {
			if i != 0 {
				sb.WriteByte(' ')
			}
			fmt.Fprintf(&sb, "%s:%d", c.sender, c.UTime)
		}
		f.log.WithField("op", "repair_one").WithField("class", f.class).
			WithField("shard", shard).WithField("uuid", id).
			WithField("msg", sb.String()).Error(err)
	}()
	return resultCh
}

func (r batchReply) UpdateTimeAt(idx int) int64 {
	if len(r.DigestData) != 0 {
		return r.DigestData[idx].UpdateTime
	}
	return r.FullData[idx].UpdateTime()
}

type _Results result[[]*storobj.Object]

func (f *reader) readAll(ctx context.Context, shard string, ids []strfmt.UUID, ch <-chan simpleResult[batchReply], st rState) <-chan _Results {
	resultCh := make(chan _Results, 1)

	go func() {
		defer close(resultCh)
		var (
			N = len(ids) // number of requested objects
			// votes counts number of votes per object for each node
			votes      = make([]vote, 0, len(st.Hosts))
			contentIdx = -1 // index of full read reply
		)

		for r := range ch { // len(ch) == st.Level
			resp := r.Response
			if r.Err != nil { // a least one node is not responding
				f.log.WithField("op", "get").WithField("replica", r.Response.Sender).
					WithField("class", f.class).WithField("shard", shard).Error(r.Err)
				resultCh <- _Results{nil, errRead}
				return
			}
			if !resp.IsDigest {
				contentIdx = len(votes)
			}

			votes = append(votes, vote{resp, make([]int, N), nil})
			M := 0
			for i := 0; i < N; i++ {
				max := 0
				lastTime := resp.UpdateTimeAt(i)

				for j := range votes { // count votes
					if votes[j].UpdateTimeAt(i) == lastTime {
						votes[j].Count[i]++
					}
					if max < votes[j].Count[i] {
						max = votes[j].Count[i]
					}
				}
				if max >= st.Level {
					M++
				}
			}

			if M == N {
				resultCh <- _Results{fromReplicas(votes[contentIdx].FullData), nil}
				return
			}
		}
		res, err := f.repairAll(ctx, shard, ids, votes, st, contentIdx)
		if err == nil {
			resultCh <- _Results{res, nil}
		}
		resultCh <- _Results{nil, errRepair}
		f.log.WithField("op", "repair_all").WithField("class", f.class).
			WithField("shard", shard).WithField("uuids", ids).Error(err)
	}()

	return resultCh
}

func (f *reader) readExistence(ctx context.Context,
	shard string,
	id strfmt.UUID,
	ch <-chan simpleResult[existReply],
	st rState,
) <-chan result[bool] {
	resultCh := make(chan result[bool], 1)
	go func() {
		defer close(resultCh)
		var (
			votes    = make([]boolTuple, 0, len(st.Hosts)) // number of votes per replica
			maxCount = 0
		)

		for r := range ch { // len(ch) == st.Level
			resp := r.Response
			if r.Err != nil { // a least one node is not responding
				f.log.WithField("op", "exists").WithField("replica", resp.Sender).
					WithField("class", f.class).WithField("shard", shard).
					WithField("uuid", id).Error(r.Err)
				resultCh <- result[bool]{false, errRead}
				return
			}

			votes = append(votes, boolTuple{resp.Sender, resp.UpdateTime, resp.RepairResponse, 0, nil})
			for i := range votes { // count number of votes
				if votes[i].UTime == resp.UpdateTime {
					votes[i].ack++
				}
				if maxCount < votes[i].ack {
					maxCount = votes[i].ack
				}
				if maxCount >= st.Level {
					exists := !votes[i].o.Deleted && votes[i].o.UpdateTime != 0
					resultCh <- result[bool]{exists, nil}
					return
				}
			}
		}

		obj, err := f.repairExist(ctx, shard, id, votes, st)
		if err == nil {
			resultCh <- result[bool]{obj, nil}
			return
		}
		resultCh <- result[bool]{false, errRepair}

		var sb strings.Builder
		for i, c := range votes {
			if i != 0 {
				sb.WriteByte(' ')
			}
			fmt.Fprintf(&sb, "%s:%d", c.sender, c.UTime)
		}
		f.log.WithField("op", "repair_exist").WithField("class", f.class).
			WithField("shard", shard).WithField("uuid", id).
			WithField("msg", sb.String()).Error(err)
	}()
	return resultCh
}
