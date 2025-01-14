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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSearchUsersParams creates a new GetSearchUsersParams object
// with the default values initialized.
func NewGetSearchUsersParams() *GetSearchUsersParams {
	var ()
	return &GetSearchUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSearchUsersParamsWithTimeout creates a new GetSearchUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSearchUsersParamsWithTimeout(timeout time.Duration) *GetSearchUsersParams {
	var ()
	return &GetSearchUsersParams{

		timeout: timeout,
	}
}

// NewGetSearchUsersParamsWithContext creates a new GetSearchUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSearchUsersParamsWithContext(ctx context.Context) *GetSearchUsersParams {
	var ()
	return &GetSearchUsersParams{

		Context: ctx,
	}
}

// NewGetSearchUsersParamsWithHTTPClient creates a new GetSearchUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSearchUsersParamsWithHTTPClient(client *http.Client) *GetSearchUsersParams {
	var ()
	return &GetSearchUsersParams{
		HTTPClient: client,
	}
}

/*GetSearchUsersParams contains all the parameters to send to the API endpoint
for the get search users operation typically these are written to a http.Request
*/
type GetSearchUsersParams struct {

	/*Limit
	  取得件数

	*/
	Limit *int64
	/*Page
	  ページ

	*/
	Page *int64
	/*Query
	  検索ワード

	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get search users params
func (o *GetSearchUsersParams) WithTimeout(timeout time.Duration) *GetSearchUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get search users params
func (o *GetSearchUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get search users params
func (o *GetSearchUsersParams) WithContext(ctx context.Context) *GetSearchUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get search users params
func (o *GetSearchUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get search users params
func (o *GetSearchUsersParams) WithHTTPClient(client *http.Client) *GetSearchUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get search users params
func (o *GetSearchUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get search users params
func (o *GetSearchUsersParams) WithLimit(limit *int64) *GetSearchUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get search users params
func (o *GetSearchUsersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get search users params
func (o *GetSearchUsersParams) WithPage(page *int64) *GetSearchUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get search users params
func (o *GetSearchUsersParams) SetPage(page *int64) {
	o.Page = page
}

// WithQuery adds the query to the get search users params
func (o *GetSearchUsersParams) WithQuery(query string) *GetSearchUsersParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get search users params
func (o *GetSearchUsersParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetSearchUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
