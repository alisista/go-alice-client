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

// NewGetMeWalletBalanceParams creates a new GetMeWalletBalanceParams object
// with the default values initialized.
func NewGetMeWalletBalanceParams() *GetMeWalletBalanceParams {

	return &GetMeWalletBalanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeWalletBalanceParamsWithTimeout creates a new GetMeWalletBalanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeWalletBalanceParamsWithTimeout(timeout time.Duration) *GetMeWalletBalanceParams {

	return &GetMeWalletBalanceParams{

		timeout: timeout,
	}
}

// NewGetMeWalletBalanceParamsWithContext creates a new GetMeWalletBalanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeWalletBalanceParamsWithContext(ctx context.Context) *GetMeWalletBalanceParams {

	return &GetMeWalletBalanceParams{

		Context: ctx,
	}
}

// NewGetMeWalletBalanceParamsWithHTTPClient creates a new GetMeWalletBalanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeWalletBalanceParamsWithHTTPClient(client *http.Client) *GetMeWalletBalanceParams {

	return &GetMeWalletBalanceParams{
		HTTPClient: client,
	}
}

/*GetMeWalletBalanceParams contains all the parameters to send to the API endpoint
for the get me wallet balance operation typically these are written to a http.Request
*/
type GetMeWalletBalanceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me wallet balance params
func (o *GetMeWalletBalanceParams) WithTimeout(timeout time.Duration) *GetMeWalletBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me wallet balance params
func (o *GetMeWalletBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me wallet balance params
func (o *GetMeWalletBalanceParams) WithContext(ctx context.Context) *GetMeWalletBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me wallet balance params
func (o *GetMeWalletBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me wallet balance params
func (o *GetMeWalletBalanceParams) WithHTTPClient(client *http.Client) *GetMeWalletBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me wallet balance params
func (o *GetMeWalletBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeWalletBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
