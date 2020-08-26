// Code generated by go-swagger; DO NOT EDIT.

package fundingrequest_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// PostTransactionFundingrequestsReader is a Reader for the PostTransactionFundingrequests structure.
type PostTransactionFundingrequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionFundingrequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionFundingrequestsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionFundingrequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostTransactionFundingrequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewPostTransactionFundingrequestsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionFundingrequestsCreated creates a PostTransactionFundingrequestsCreated with default headers values
func NewPostTransactionFundingrequestsCreated() *PostTransactionFundingrequestsCreated {
	return &PostTransactionFundingrequestsCreated{}
}

/*PostTransactionFundingrequestsCreated handles this case with default header values.

creation response
*/
type PostTransactionFundingrequestsCreated struct {

	//Payload

	// isStream: false
	*models.FundingRequestResponse
}

func (o *PostTransactionFundingrequestsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/fundingrequests][%d] postTransactionFundingrequestsCreated", 201)
}

func (o *PostTransactionFundingrequestsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FundingRequestResponse = new(models.FundingRequestResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FundingRequestResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionFundingrequestsBadRequest creates a PostTransactionFundingrequestsBadRequest with default headers values
func NewPostTransactionFundingrequestsBadRequest() *PostTransactionFundingrequestsBadRequest {
	return &PostTransactionFundingrequestsBadRequest{}
}

/*PostTransactionFundingrequestsBadRequest handles this case with default header values.

bad request
*/
type PostTransactionFundingrequestsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionFundingrequestsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/fundingrequests][%d] postTransactionFundingrequestsBadRequest", 400)
}

func (o *PostTransactionFundingrequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionFundingrequestsForbidden creates a PostTransactionFundingrequestsForbidden with default headers values
func NewPostTransactionFundingrequestsForbidden() *PostTransactionFundingrequestsForbidden {
	return &PostTransactionFundingrequestsForbidden{}
}

/*PostTransactionFundingrequestsForbidden handles this case with default header values.

forbidden
*/
type PostTransactionFundingrequestsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionFundingrequestsForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/fundingrequests][%d] postTransactionFundingrequestsForbidden", 403)
}

func (o *PostTransactionFundingrequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionFundingrequestsBadGateway creates a PostTransactionFundingrequestsBadGateway with default headers values
func NewPostTransactionFundingrequestsBadGateway() *PostTransactionFundingrequestsBadGateway {
	return &PostTransactionFundingrequestsBadGateway{}
}

/*PostTransactionFundingrequestsBadGateway handles this case with default header values.

gateway issue
*/
type PostTransactionFundingrequestsBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionFundingrequestsBadGateway) Error() string {
	return fmt.Sprintf("[POST /transaction/fundingrequests][%d] postTransactionFundingrequestsBadGateway", 502)
}

func (o *PostTransactionFundingrequestsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
