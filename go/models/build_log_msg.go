// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BuildLogMsg build log msg
// swagger:model buildLogMsg

type BuildLogMsg struct {

	// error
	Error bool `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

/* polymorph buildLogMsg error false */

/* polymorph buildLogMsg message false */

// Validate validates this build log msg
func (m *BuildLogMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BuildLogMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildLogMsg) UnmarshalBinary(b []byte) error {
	var res BuildLogMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
