package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/opensds/opensds/api/openapi-spec/models"
)

// UpdateShareReader is a Reader for the UpdateShare structure.
type UpdateShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateShareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateShareOK creates a UpdateShareOK with default headers values
func NewUpdateShareOK() *UpdateShareOK {
	return &UpdateShareOK{}
}

/*UpdateShareOK handles this case with default header values.

OK
*/
type UpdateShareOK struct {
	Payload *models.ShareResponse
}

func (o *UpdateShareOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/shares/{resourceType}/{id}][%d] updateShareOK  %+v", 200, o.Payload)
}

func (o *UpdateShareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShareResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}