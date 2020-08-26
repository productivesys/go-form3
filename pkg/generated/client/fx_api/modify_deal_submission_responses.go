// Code generated by go-swagger; DO NOT EDIT.

package fx_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// ModifyDealSubmissionReader is a Reader for the ModifyDealSubmission structure.
type ModifyDealSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyDealSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyDealSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyDealSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewModifyDealSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewModifyDealSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewModifyDealSubmissionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyDealSubmissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyDealSubmissionOK creates a ModifyDealSubmissionOK with default headers values
func NewModifyDealSubmissionOK() *ModifyDealSubmissionOK {
	return &ModifyDealSubmissionOK{}
}

/*ModifyDealSubmissionOK handles this case with default header values.

creation response
*/
type ModifyDealSubmissionOK struct {

	//Payload

	// isStream: false
	*models.FXDealSubmissionResponse
}

func (o *ModifyDealSubmissionOK) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] modifyDealSubmissionOK", 200)
}

func (o *ModifyDealSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FXDealSubmissionResponse = new(models.FXDealSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FXDealSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDealSubmissionBadRequest creates a ModifyDealSubmissionBadRequest with default headers values
func NewModifyDealSubmissionBadRequest() *ModifyDealSubmissionBadRequest {
	return &ModifyDealSubmissionBadRequest{}
}

/*ModifyDealSubmissionBadRequest handles this case with default header values.

Bad Request
*/
type ModifyDealSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifyDealSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] modifyDealSubmissionBadRequest", 400)
}

func (o *ModifyDealSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDealSubmissionForbidden creates a ModifyDealSubmissionForbidden with default headers values
func NewModifyDealSubmissionForbidden() *ModifyDealSubmissionForbidden {
	return &ModifyDealSubmissionForbidden{}
}

/*ModifyDealSubmissionForbidden handles this case with default header values.

Action Forbidden
*/
type ModifyDealSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifyDealSubmissionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] modifyDealSubmissionForbidden", 403)
}

func (o *ModifyDealSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDealSubmissionNotFound creates a ModifyDealSubmissionNotFound with default headers values
func NewModifyDealSubmissionNotFound() *ModifyDealSubmissionNotFound {
	return &ModifyDealSubmissionNotFound{}
}

/*ModifyDealSubmissionNotFound handles this case with default header values.

Not Found
*/
type ModifyDealSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifyDealSubmissionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] modifyDealSubmissionNotFound", 404)
}

func (o *ModifyDealSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDealSubmissionBadGateway creates a ModifyDealSubmissionBadGateway with default headers values
func NewModifyDealSubmissionBadGateway() *ModifyDealSubmissionBadGateway {
	return &ModifyDealSubmissionBadGateway{}
}

/*ModifyDealSubmissionBadGateway handles this case with default header values.

Bad Gateway
*/
type ModifyDealSubmissionBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifyDealSubmissionBadGateway) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] modifyDealSubmissionBadGateway", 502)
}

func (o *ModifyDealSubmissionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyDealSubmissionDefault creates a ModifyDealSubmissionDefault with default headers values
func NewModifyDealSubmissionDefault(code int) *ModifyDealSubmissionDefault {
	return &ModifyDealSubmissionDefault{
		_statusCode: code,
	}
}

/*ModifyDealSubmissionDefault handles this case with default header values.

Unexpected Error
*/
type ModifyDealSubmissionDefault struct {
	_statusCode int

	//Payload

	// isStream: false
	*models.APIError
}

// Code gets the status code for the modify deal submission default response
func (o *ModifyDealSubmissionDefault) Code() int {
	return o._statusCode
}

func (o *ModifyDealSubmissionDefault) Error() string {
	return fmt.Sprintf("[PATCH /fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}][%d] ModifyDealSubmission default", o._statusCode)
}

func (o *ModifyDealSubmissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
