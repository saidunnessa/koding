package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJAccountFetchOAuthInfoIDParams creates a new PostRemoteAPIJAccountFetchOAuthInfoIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountFetchOAuthInfoIDParams() *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchOAuthInfoIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountFetchOAuthInfoIDParamsWithTimeout creates a new PostRemoteAPIJAccountFetchOAuthInfoIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountFetchOAuthInfoIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchOAuthInfoIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountFetchOAuthInfoIDParamsWithContext creates a new PostRemoteAPIJAccountFetchOAuthInfoIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountFetchOAuthInfoIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchOAuthInfoIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountFetchOAuthInfoIDParams contains all the parameters to send to the API endpoint
for the post remote API j account fetch o auth info ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountFetchOAuthInfoIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) WithID(id string) *PostRemoteAPIJAccountFetchOAuthInfoIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account fetch o auth info ID params
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountFetchOAuthInfoIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
