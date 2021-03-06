package j_group

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

// NewJGroupFetchBlockedAccountsParams creates a new JGroupFetchBlockedAccountsParams object
// with the default values initialized.
func NewJGroupFetchBlockedAccountsParams() *JGroupFetchBlockedAccountsParams {
	var ()
	return &JGroupFetchBlockedAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchBlockedAccountsParamsWithTimeout creates a new JGroupFetchBlockedAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchBlockedAccountsParamsWithTimeout(timeout time.Duration) *JGroupFetchBlockedAccountsParams {
	var ()
	return &JGroupFetchBlockedAccountsParams{

		timeout: timeout,
	}
}

// NewJGroupFetchBlockedAccountsParamsWithContext creates a new JGroupFetchBlockedAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchBlockedAccountsParamsWithContext(ctx context.Context) *JGroupFetchBlockedAccountsParams {
	var ()
	return &JGroupFetchBlockedAccountsParams{

		Context: ctx,
	}
}

/*JGroupFetchBlockedAccountsParams contains all the parameters to send to the API endpoint
for the j group fetch blocked accounts operation typically these are written to a http.Request
*/
type JGroupFetchBlockedAccountsParams struct {

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

// WithTimeout adds the timeout to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) WithTimeout(timeout time.Duration) *JGroupFetchBlockedAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) WithContext(ctx context.Context) *JGroupFetchBlockedAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) WithBody(body models.DefaultSelector) *JGroupFetchBlockedAccountsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) WithID(id string) *JGroupFetchBlockedAccountsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch blocked accounts params
func (o *JGroupFetchBlockedAccountsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchBlockedAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
