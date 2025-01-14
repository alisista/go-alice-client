// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Notification notification
// swagger:model Notification
type Notification struct {

	// acted user id
	ActedUserID string `json:"acted_user_id,omitempty"`

	// created at
	CreatedAt float64 `json:"created_at,omitempty"`

	// sort key
	SortKey int64 `json:"sort_key,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
