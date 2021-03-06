package j_reward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JRewardAddCustomRewardReader is a Reader for the JRewardAddCustomReward structure.
type JRewardAddCustomRewardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JRewardAddCustomRewardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJRewardAddCustomRewardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJRewardAddCustomRewardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJRewardAddCustomRewardOK creates a JRewardAddCustomRewardOK with default headers values
func NewJRewardAddCustomRewardOK() *JRewardAddCustomRewardOK {
	return &JRewardAddCustomRewardOK{}
}

/*JRewardAddCustomRewardOK handles this case with default header values.

Request processed successfully
*/
type JRewardAddCustomRewardOK struct {
	Payload *models.DefaultResponse
}

func (o *JRewardAddCustomRewardOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.addCustomReward][%d] jRewardAddCustomRewardOK  %+v", 200, o.Payload)
}

func (o *JRewardAddCustomRewardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJRewardAddCustomRewardUnauthorized creates a JRewardAddCustomRewardUnauthorized with default headers values
func NewJRewardAddCustomRewardUnauthorized() *JRewardAddCustomRewardUnauthorized {
	return &JRewardAddCustomRewardUnauthorized{}
}

/*JRewardAddCustomRewardUnauthorized handles this case with default header values.

Unauthorized request
*/
type JRewardAddCustomRewardUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JRewardAddCustomRewardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.addCustomReward][%d] jRewardAddCustomRewardUnauthorized  %+v", 401, o.Payload)
}

func (o *JRewardAddCustomRewardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
