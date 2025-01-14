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

	models "github.com/alisista/go-alis-client/models"
)

// NewPutMeArticlesArticleIDDraftsPublishParams creates a new PutMeArticlesArticleIDDraftsPublishParams object
// with the default values initialized.
func NewPutMeArticlesArticleIDDraftsPublishParams() *PutMeArticlesArticleIDDraftsPublishParams {
	var ()
	return &PutMeArticlesArticleIDDraftsPublishParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeArticlesArticleIDDraftsPublishParamsWithTimeout creates a new PutMeArticlesArticleIDDraftsPublishParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeArticlesArticleIDDraftsPublishParamsWithTimeout(timeout time.Duration) *PutMeArticlesArticleIDDraftsPublishParams {
	var ()
	return &PutMeArticlesArticleIDDraftsPublishParams{

		timeout: timeout,
	}
}

// NewPutMeArticlesArticleIDDraftsPublishParamsWithContext creates a new PutMeArticlesArticleIDDraftsPublishParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeArticlesArticleIDDraftsPublishParamsWithContext(ctx context.Context) *PutMeArticlesArticleIDDraftsPublishParams {
	var ()
	return &PutMeArticlesArticleIDDraftsPublishParams{

		Context: ctx,
	}
}

// NewPutMeArticlesArticleIDDraftsPublishParamsWithHTTPClient creates a new PutMeArticlesArticleIDDraftsPublishParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeArticlesArticleIDDraftsPublishParamsWithHTTPClient(client *http.Client) *PutMeArticlesArticleIDDraftsPublishParams {
	var ()
	return &PutMeArticlesArticleIDDraftsPublishParams{
		HTTPClient: client,
	}
}

/*PutMeArticlesArticleIDDraftsPublishParams contains all the parameters to send to the API endpoint
for the put me articles article ID drafts publish operation typically these are written to a http.Request
*/
type PutMeArticlesArticleIDDraftsPublishParams struct {

	/*ArticleID
	  対象記事の指定するために使用

	*/
	ArticleID string
	/*Publish
	  publish object

	*/
	Publish *models.Publish

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) WithTimeout(timeout time.Duration) *PutMeArticlesArticleIDDraftsPublishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) WithContext(ctx context.Context) *PutMeArticlesArticleIDDraftsPublishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) WithHTTPClient(client *http.Client) *PutMeArticlesArticleIDDraftsPublishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleID adds the articleID to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) WithArticleID(articleID string) *PutMeArticlesArticleIDDraftsPublishParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WithPublish adds the publish to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) WithPublish(publish *models.Publish) *PutMeArticlesArticleIDDraftsPublishParams {
	o.SetPublish(publish)
	return o
}

// SetPublish adds the publish to the put me articles article ID drafts publish params
func (o *PutMeArticlesArticleIDDraftsPublishParams) SetPublish(publish *models.Publish) {
	o.Publish = publish
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeArticlesArticleIDDraftsPublishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param article_id
	if err := r.SetPathParam("article_id", o.ArticleID); err != nil {
		return err
	}

	if o.Publish != nil {
		if err := r.SetBodyParam(o.Publish); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
