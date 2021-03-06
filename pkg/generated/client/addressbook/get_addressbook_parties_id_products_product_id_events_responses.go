// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetAddressbookPartiesIDProductsProductIDEventsReader is a Reader for the GetAddressbookPartiesIDProductsProductIDEvents structure.
type GetAddressbookPartiesIDProductsProductIDEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDProductsProductIDEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDProductsProductIDEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDProductsProductIDEventsOK creates a GetAddressbookPartiesIDProductsProductIDEventsOK with default headers values
func NewGetAddressbookPartiesIDProductsProductIDEventsOK() *GetAddressbookPartiesIDProductsProductIDEventsOK {
	return &GetAddressbookPartiesIDProductsProductIDEventsOK{}
}

/*GetAddressbookPartiesIDProductsProductIDEventsOK handles this case with default header values.

list contact accounts response
*/
type GetAddressbookPartiesIDProductsProductIDEventsOK struct {

	//Payload

	// isStream: false
	*models.ListPartyProductEventsResponse
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}/events][%d] getAddressbookPartiesIdProductsProductIdEventsOK", 200)
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ListPartyProductEventsResponse = new(models.ListPartyProductEventsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ListPartyProductEventsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
