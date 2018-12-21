// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMeInfoParams creates a new GetMeInfoParams object
// with the default values initialized.
func NewGetMeInfoParams() *GetMeInfoParams {

	return &GetMeInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeInfoParamsWithTimeout creates a new GetMeInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeInfoParamsWithTimeout(timeout time.Duration) *GetMeInfoParams {

	return &GetMeInfoParams{

		timeout: timeout,
	}
}

// NewGetMeInfoParamsWithContext creates a new GetMeInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeInfoParamsWithContext(ctx context.Context) *GetMeInfoParams {

	return &GetMeInfoParams{

		Context: ctx,
	}
}

// NewGetMeInfoParamsWithHTTPClient creates a new GetMeInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeInfoParamsWithHTTPClient(client *http.Client) *GetMeInfoParams {

	return &GetMeInfoParams{
		HTTPClient: client,
	}
}

/*GetMeInfoParams contains all the parameters to send to the API endpoint
for the get me info operation typically these are written to a http.Request
*/
type GetMeInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me info params
func (o *GetMeInfoParams) WithTimeout(timeout time.Duration) *GetMeInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me info params
func (o *GetMeInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me info params
func (o *GetMeInfoParams) WithContext(ctx context.Context) *GetMeInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me info params
func (o *GetMeInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me info params
func (o *GetMeInfoParams) WithHTTPClient(client *http.Client) *GetMeInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me info params
func (o *GetMeInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}