// Code generated by go-swagger; DO NOT EDIT.

package fundingrequest_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetTransactionFundingrequestsReader is a Reader for the GetTransactionFundingrequests structure.
type GetTransactionFundingrequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFundingrequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFundingrequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionFundingrequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetTransactionFundingrequestsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFundingrequestsOK creates a GetTransactionFundingrequestsOK with default headers values
func NewGetTransactionFundingrequestsOK() *GetTransactionFundingrequestsOK {
	return &GetTransactionFundingrequestsOK{}
}

/*GetTransactionFundingrequestsOK handles this case with default header values.

Funding Requests response
*/
type GetTransactionFundingrequestsOK struct {

	//Payload

	// isStream: false
	*models.FundingRequestListResponse
}

func (o *GetTransactionFundingrequestsOK) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests][%d] getTransactionFundingrequestsOK", 200)
}

func (o *GetTransactionFundingrequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FundingRequestListResponse = new(models.FundingRequestListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FundingRequestListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFundingrequestsBadRequest creates a GetTransactionFundingrequestsBadRequest with default headers values
func NewGetTransactionFundingrequestsBadRequest() *GetTransactionFundingrequestsBadRequest {
	return &GetTransactionFundingrequestsBadRequest{}
}

/*GetTransactionFundingrequestsBadRequest handles this case with default header values.

Funding Requests bad request
*/
type GetTransactionFundingrequestsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFundingrequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests][%d] getTransactionFundingrequestsBadRequest", 400)
}

func (o *GetTransactionFundingrequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFundingrequestsBadGateway creates a GetTransactionFundingrequestsBadGateway with default headers values
func NewGetTransactionFundingrequestsBadGateway() *GetTransactionFundingrequestsBadGateway {
	return &GetTransactionFundingrequestsBadGateway{}
}

/*GetTransactionFundingrequestsBadGateway handles this case with default header values.

Funding Requests gateway issue
*/
type GetTransactionFundingrequestsBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFundingrequestsBadGateway) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests][%d] getTransactionFundingrequestsBadGateway", 502)
}

func (o *GetTransactionFundingrequestsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
