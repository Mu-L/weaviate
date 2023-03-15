package cluster

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskSpace(t *testing.T) {
	for _, name := range []string{"", "host-12:1", "2", "00", "-jhd"} {
		want := nodeSpace{
			name,
			DiskSpace{
				Total:     256,
				Available: 3,
			},
		}
		bytes, err := want.marshal()
		assert.Nil(t, err)
		got := nodeSpace{}
		err = got.Unmarshal(bytes)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestDelegate(t *testing.T) {
	st := State{
		delegate: delegate{
			dataPath:  ".",
			DiskUsage: make(map[string]DiskSpace, 32),
		},
	}
	diskSpaces := []DiskSpace{{1, 1}, {3, 2}, {4, 5}, {4, 2}, {4, 1}, {5, 2}}
	done := make(chan struct{})
	go func() {
		for i, x := range diskSpaces {
			node := fmt.Sprintf("N-%d", i+1)
			st.delegate.Set(node, x)
		}
		done <- struct{}{}
	}()

	_, ok := st.delegate.Get("X")
	assert.False(t, ok)

	for i, x := range diskSpaces {
		space, ok := st.delegate.Get(fmt.Sprintf("N-%d", i+1))
		if ok {
			assert.Equal(t, x, space)
		}
	}
	<-done
	for i, x := range diskSpaces {
		node := fmt.Sprintf("N-%d", i+1)
		space, ok := st.delegate.Get(node)
		assert.Equal(t, x, space)
		assert.True(t, ok)
		st.delegate.Delete(node)

	}
	assert.Empty(t, st.delegate.DiskUsage)
}
