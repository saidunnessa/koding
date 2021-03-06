package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialChannelLeaveReader is a Reader for the SocialChannelLeave structure.
type SocialChannelLeaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialChannelLeaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialChannelLeaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialChannelLeaveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialChannelLeaveOK creates a SocialChannelLeaveOK with default headers values
func NewSocialChannelLeaveOK() *SocialChannelLeaveOK {
	return &SocialChannelLeaveOK{}
}

/*SocialChannelLeaveOK handles this case with default header values.

Request processed successfully
*/
type SocialChannelLeaveOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialChannelLeaveOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.leave][%d] socialChannelLeaveOK  %+v", 200, o.Payload)
}

func (o *SocialChannelLeaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialChannelLeaveUnauthorized creates a SocialChannelLeaveUnauthorized with default headers values
func NewSocialChannelLeaveUnauthorized() *SocialChannelLeaveUnauthorized {
	return &SocialChannelLeaveUnauthorized{}
}

/*SocialChannelLeaveUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialChannelLeaveUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialChannelLeaveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.leave][%d] socialChannelLeaveUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialChannelLeaveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
