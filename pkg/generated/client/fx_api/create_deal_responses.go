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

// CreateDealReader is a Reader for the CreateDeal structure.
type CreateDealReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDealReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDealCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDealBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateDealForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewCreateDealBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDealCreated creates a CreateDealCreated with default headers values
func NewCreateDealCreated() *CreateDealCreated {
	return &CreateDealCreated{}
}

/*CreateDealCreated handles this case with default header values.

creation response
*/
type CreateDealCreated struct {

	//Payload

	// isStream: false
	*models.FXDealResponse
}

func (o *CreateDealCreated) Error() string {
	return fmt.Sprintf("[POST /fx/deals][%d] createDealCreated", 201)
}

func (o *CreateDealCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FXDealResponse = new(models.FXDealResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FXDealResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDealBadRequest creates a CreateDealBadRequest with default headers values
func NewCreateDealBadRequest() *CreateDealBadRequest {
	return &CreateDealBadRequest{}
}

/*CreateDealBadRequest handles this case with default header values.

bad request
*/
type CreateDealBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDealBadRequest) Error() string {
	return fmt.Sprintf("[POST /fx/deals][%d] createDealBadRequest", 400)
}

func (o *CreateDealBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDealForbidden creates a CreateDealForbidden with default headers values
func NewCreateDealForbidden() *CreateDealForbidden {
	return &CreateDealForbidden{}
}

/*CreateDealForbidden handles this case with default header values.

forbidden
*/
type CreateDealForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDealForbidden) Error() string {
	return fmt.Sprintf("[POST /fx/deals][%d] createDealForbidden", 403)
}

func (o *CreateDealForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDealBadGateway creates a CreateDealBadGateway with default headers values
func NewCreateDealBadGateway() *CreateDealBadGateway {
	return &CreateDealBadGateway{}
}

/*CreateDealBadGateway handles this case with default header values.

FX rates gateway issue
*/
type CreateDealBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDealBadGateway) Error() string {
	return fmt.Sprintf("[POST /fx/deals][%d] createDealBadGateway", 502)
}

func (o *CreateDealBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}