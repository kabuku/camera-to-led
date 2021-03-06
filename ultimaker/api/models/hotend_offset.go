package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HotendOffset X,Y and Z offset of this hotend nozzle exit compared to the other hotends in this head. The state indicates if the data for this hotend is valid and thus can be used.
// swagger:model HotendOffset
type HotendOffset struct {

	// state
	State string `json:"state,omitempty"`

	// x
	X float64 `json:"x,omitempty"`

	// y
	Y float64 `json:"y,omitempty"`

	// z
	Z float64 `json:"z,omitempty"`
}

// Validate validates this hotend offset
func (m *HotendOffset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hotendOffsetTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hotendOffsetTypeStatePropEnum = append(hotendOffsetTypeStatePropEnum, v)
	}
}

const (
	// HotendOffsetStateValid captures enum value "valid"
	HotendOffsetStateValid string = "valid"
	// HotendOffsetStateInvalid captures enum value "invalid"
	HotendOffsetStateInvalid string = "invalid"
)

// prop value enum
func (m *HotendOffset) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hotendOffsetTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HotendOffset) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}
