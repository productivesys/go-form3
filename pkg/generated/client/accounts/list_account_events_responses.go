// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// ListAccountEventsReader is a Reader for the ListAccountEvents structure.
type ListAccountEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountEventsOK creates a ListAccountEventsOK with default headers values
func NewListAccountEventsOK() *ListAccountEventsOK {
	return &ListAccountEventsOK{}
}

/*ListAccountEventsOK handles this case with default header values.

Account event list
*/
type ListAccountEventsOK struct {

	//Payload

	// isStream: false
	*models.AccountEventListResponse
}

func (o *ListAccountEventsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}/events][%d] listAccountEventsOK", 200)
}

func (o *ListAccountEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountEventListResponse = new(models.AccountEventListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountEventListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
