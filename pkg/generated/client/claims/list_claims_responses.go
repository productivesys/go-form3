// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// ListClaimsReader is a Reader for the ListClaims structure.
type ListClaimsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClaimsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClaimsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClaimsOK creates a ListClaimsOK with default headers values
func NewListClaimsOK() *ListClaimsOK {
	return &ListClaimsOK{}
}

/*ListClaimsOK handles this case with default header values.

List of claims details
*/
type ListClaimsOK struct {

	//Payload

	// isStream: false
	*models.ClaimDetailsListResponse
}

func (o *ListClaimsOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims][%d] listClaimsOK", 200)
}

func (o *ListClaimsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimDetailsListResponse = new(models.ClaimDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
