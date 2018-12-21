// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// PostMeArticlesArticleIDImagesReader is a Reader for the PostMeArticlesArticleIDImages structure.
type PostMeArticlesArticleIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeArticlesArticleIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeArticlesArticleIDImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostMeArticlesArticleIDImagesOK creates a PostMeArticlesArticleIDImagesOK with default headers values
func NewPostMeArticlesArticleIDImagesOK() *PostMeArticlesArticleIDImagesOK {
	return &PostMeArticlesArticleIDImagesOK{}
}

/*PostMeArticlesArticleIDImagesOK handles this case with default header values.

登録した画像データのURL
*/
type PostMeArticlesArticleIDImagesOK struct {
	Payload *PostMeArticlesArticleIDImagesOKBody
}

func (o *PostMeArticlesArticleIDImagesOK) Error() string {
	return fmt.Sprintf("[POST /me/articles/{article_id}/images][%d] postMeArticlesArticleIdImagesOK  %+v", 200, o.Payload)
}

func (o *PostMeArticlesArticleIDImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostMeArticlesArticleIDImagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostMeArticlesArticleIDImagesOKBody post me articles article ID images o k body
swagger:model PostMeArticlesArticleIDImagesOKBody
*/
type PostMeArticlesArticleIDImagesOKBody struct {

	// image url
	ImageURL string `json:"image_url,omitempty"`
}

// Validate validates this post me articles article ID images o k body
func (o *PostMeArticlesArticleIDImagesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMeArticlesArticleIDImagesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMeArticlesArticleIDImagesOKBody) UnmarshalBinary(b []byte) error {
	var res PostMeArticlesArticleIDImagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}