// Code generated by go-swagger; DO NOT EDIT.

package fx_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetDealReader is a Reader for the GetDeal structure.
type GetDealReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDealReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDealOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDealBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDealForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDealOK creates a GetDealOK with default headers values
func NewGetDealOK() *GetDealOK {
	return &GetDealOK{}
}

/*GetDealOK handles this case with default header values.

fx deal response
*/
type GetDealOK struct {

	//Payload

	// isStream: false
	*models.FXDealResponse
}

func (o *GetDealOK) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealOK", 200)
}

func (o *GetDealOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FXDealResponse = new(models.FXDealResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FXDealResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealBadRequest creates a GetDealBadRequest with default headers values
func NewGetDealBadRequest() *GetDealBadRequest {
	return &GetDealBadRequest{}
}

/*GetDealBadRequest handles this case with default header values.

bad request
*/
type GetDealBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealBadRequest) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealBadRequest", 400)
}

func (o *GetDealBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealForbidden creates a GetDealForbidden with default headers values
func NewGetDealForbidden() *GetDealForbidden {
	return &GetDealForbidden{}
}

/*GetDealForbidden handles this case with default header values.

forbidden
*/
type GetDealForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealForbidden) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealForbidden", 403)
}

func (o *GetDealForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
