package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// XYZ Container for xyz
// swagger:model XYZ
type XYZ struct {

	// x
	X float64 `json:"x,omitempty"`

	// y
	Y float64 `json:"y,omitempty"`

	// z
	Z float64 `json:"z,omitempty"`
}

// Validate validates this x y z
func (m *XYZ) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}