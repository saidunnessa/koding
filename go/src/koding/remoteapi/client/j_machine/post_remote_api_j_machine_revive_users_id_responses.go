package j_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJMachineReviveUsersIDReader is a Reader for the PostRemoteAPIJMachineReviveUsersID structure.
type PostRemoteAPIJMachineReviveUsersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJMachineReviveUsersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJMachineReviveUsersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJMachineReviveUsersIDOK creates a PostRemoteAPIJMachineReviveUsersIDOK with default headers values
func NewPostRemoteAPIJMachineReviveUsersIDOK() *PostRemoteAPIJMachineReviveUsersIDOK {
	return &PostRemoteAPIJMachineReviveUsersIDOK{}
}

/*PostRemoteAPIJMachineReviveUsersIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJMachineReviveUsersIDOK struct {
	Payload PostRemoteAPIJMachineReviveUsersIDOKBody
}

func (o *PostRemoteAPIJMachineReviveUsersIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.reviveUsers/{id}][%d] postRemoteApiJMachineReviveUsersIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJMachineReviveUsersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJMachineReviveUsersIDOKBody post remote API j machine revive users ID o k body
swagger:model PostRemoteAPIJMachineReviveUsersIDOKBody
*/
type PostRemoteAPIJMachineReviveUsersIDOKBody struct {
	models.JMachine

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJMachineReviveUsersIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJMachineReviveUsersIDOKBodyAO0 models.JMachine
	if err := swag.ReadJSON(raw, &postRemoteAPIJMachineReviveUsersIDOKBodyAO0); err != nil {
		return err
	}
	o.JMachine = postRemoteAPIJMachineReviveUsersIDOKBodyAO0

	var postRemoteAPIJMachineReviveUsersIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJMachineReviveUsersIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJMachineReviveUsersIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJMachineReviveUsersIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJMachineReviveUsersIDOKBodyAO0, err := swag.WriteJSON(o.JMachine)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJMachineReviveUsersIDOKBodyAO0)

	postRemoteAPIJMachineReviveUsersIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJMachineReviveUsersIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j machine revive users ID o k body
func (o *PostRemoteAPIJMachineReviveUsersIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JMachine.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
