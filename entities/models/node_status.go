//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeStatus The definition of a backup node status response body
//
// swagger:model NodeStatus
type NodeStatus struct {

	// The gitHash of Weaviate.
	GitHash string `json:"gitHash,omitempty"`

	// The name of the node.
	Name string `json:"name,omitempty"`

	// The list of the shards with it's statistics.
	Shards []*NodeShardStatus `json:"shards"`

	// Weaviate overall statistics.
	Stats *NodeStats `json:"stats,omitempty"`

	// Node's status.
	// Enum: [HEALTHY UNHEALTHY UNAVAILABLE]
	Status *string `json:"status,omitempty"`

	// The version of Weaviate.
	Version string `json:"version,omitempty"`
}

// Validate validates this node status
func (m *NodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeStatus) validateShards(formats strfmt.Registry) error {

	if swag.IsZero(m.Shards) { // not required
		return nil
	}

	for i := 0; i < len(m.Shards); i++ {
		if swag.IsZero(m.Shards[i]) { // not required
			continue
		}

		if m.Shards[i] != nil {
			if err := m.Shards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NodeStatus) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var nodeStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEALTHY","UNHEALTHY","UNAVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeStatusTypeStatusPropEnum = append(nodeStatusTypeStatusPropEnum, v)
	}
}

const (

	// NodeStatusStatusHEALTHY captures enum value "HEALTHY"
	NodeStatusStatusHEALTHY string = "HEALTHY"

	// NodeStatusStatusUNHEALTHY captures enum value "UNHEALTHY"
	NodeStatusStatusUNHEALTHY string = "UNHEALTHY"

	// NodeStatusStatusUNAVAILABLE captures enum value "UNAVAILABLE"
	NodeStatusStatusUNAVAILABLE string = "UNAVAILABLE"
)

// prop value enum
func (m *NodeStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatus) UnmarshalBinary(b []byte) error {
	var res NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
