package j_account

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

// PostRemoteAPIJAccountFetchMySessionsIDReader is a Reader for the PostRemoteAPIJAccountFetchMySessionsID structure.
type PostRemoteAPIJAccountFetchMySessionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountFetchMySessionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountFetchMySessionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountFetchMySessionsIDOK creates a PostRemoteAPIJAccountFetchMySessionsIDOK with default headers values
func NewPostRemoteAPIJAccountFetchMySessionsIDOK() *PostRemoteAPIJAccountFetchMySessionsIDOK {
	return &PostRemoteAPIJAccountFetchMySessionsIDOK{}
}

/*PostRemoteAPIJAccountFetchMySessionsIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountFetchMySessionsIDOK struct {
	Payload PostRemoteAPIJAccountFetchMySessionsIDOKBody
}

func (o *PostRemoteAPIJAccountFetchMySessionsIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchMySessions/{id}][%d] postRemoteApiJAccountFetchMySessionsIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountFetchMySessionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountFetchMySessionsIDOKBody post remote API j account fetch my sessions ID o k body
swagger:model PostRemoteAPIJAccountFetchMySessionsIDOKBody
*/
type PostRemoteAPIJAccountFetchMySessionsIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountFetchMySessionsIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountFetchMySessionsIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchMySessionsIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountFetchMySessionsIDOKBodyAO0

	var postRemoteAPIJAccountFetchMySessionsIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchMySessionsIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountFetchMySessionsIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountFetchMySessionsIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountFetchMySessionsIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchMySessionsIDOKBodyAO0)

	postRemoteAPIJAccountFetchMySessionsIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchMySessionsIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account fetch my sessions ID o k body
func (o *PostRemoteAPIJAccountFetchMySessionsIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
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
