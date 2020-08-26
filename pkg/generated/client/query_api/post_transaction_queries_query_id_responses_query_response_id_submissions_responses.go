// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsReader is a Reader for the PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissions structure.
type PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated creates a PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated with default headers values
func NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated() *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated {
	return &PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated{}
}

/*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated handles this case with default header values.

creation response
*/
type PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated struct {

	//Payload

	// isStream: false
	*models.QueryResponseSubmissionResponse
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] postTransactionQueriesQueryIdResponsesQueryResponseIdSubmissionsCreated", 201)
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryResponseSubmissionResponse = new(models.QueryResponseSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryResponseSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest creates a PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest with default headers values
func NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest() *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest {
	return &PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest{}
}

/*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest handles this case with default header values.

bad request
*/
type PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] postTransactionQueriesQueryIdResponsesQueryResponseIdSubmissionsBadRequest", 400)
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden creates a PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden with default headers values
func NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden() *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden {
	return &PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden{}
}

/*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden handles this case with default header values.

forbidden
*/
type PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] postTransactionQueriesQueryIdResponsesQueryResponseIdSubmissionsForbidden", 403)
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway creates a PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway with default headers values
func NewPostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway() *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway {
	return &PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway{}
}

/*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway handles this case with default header values.

gateway issue
*/
type PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] postTransactionQueriesQueryIdResponsesQueryResponseIdSubmissionsBadGateway", 502)
}

func (o *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
