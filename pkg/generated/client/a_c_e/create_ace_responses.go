// Code generated by go-swagger; DO NOT EDIT.

package a_c_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreateAceReader is a Reader for the CreateAce structure.
type CreateAceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAceCreated creates a CreateAceCreated with default headers values
func NewCreateAceCreated() *CreateAceCreated {
	return &CreateAceCreated{}
}

/*CreateAceCreated handles this case with default header values.

ACE creation response
*/
type CreateAceCreated struct {

	//Payload

	// isStream: false
	*models.AceCreationResponse
}

func (o *CreateAceCreated) Error() string {
	return fmt.Sprintf("[POST /security/roles/{role_id}/aces][%d] createAceCreated", 201)
}

func (o *CreateAceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AceCreationResponse = new(models.AceCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AceCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
