// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ArticleContent article content
// swagger:model ArticleContent
type ArticleContent struct {

	// article id
	ArticleID string `json:"article_id,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// eye catch url
	EyeCatchURL string `json:"eye_catch_url,omitempty"`

	// overview
	Overview string `json:"overview,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this article content
func (m *ArticleContent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArticleContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArticleContent) UnmarshalBinary(b []byte) error {
	var res ArticleContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}