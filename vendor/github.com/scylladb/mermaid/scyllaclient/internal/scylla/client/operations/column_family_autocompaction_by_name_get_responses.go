// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyAutocompactionByNameGetReader is a Reader for the ColumnFamilyAutocompactionByNameGet structure.
type ColumnFamilyAutocompactionByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyAutocompactionByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyAutocompactionByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyAutocompactionByNameGetOK creates a ColumnFamilyAutocompactionByNameGetOK with default headers values
func NewColumnFamilyAutocompactionByNameGetOK() *ColumnFamilyAutocompactionByNameGetOK {
	return &ColumnFamilyAutocompactionByNameGetOK{}
}

/*ColumnFamilyAutocompactionByNameGetOK handles this case with default header values.

ColumnFamilyAutocompactionByNameGetOK column family autocompaction by name get o k
*/
type ColumnFamilyAutocompactionByNameGetOK struct {
	Payload bool
}

func (o *ColumnFamilyAutocompactionByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/autocompaction/{name}][%d] columnFamilyAutocompactionByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyAutocompactionByNameGetOK) GetPayload() bool {
	return o.Payload
}

func (o *ColumnFamilyAutocompactionByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}