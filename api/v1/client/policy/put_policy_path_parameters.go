package policy

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
)

// NewPutPolicyPathParams creates a new PutPolicyPathParams object
// with the default values initialized.
func NewPutPolicyPathParams() *PutPolicyPathParams {
	var ()
	return &PutPolicyPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPolicyPathParamsWithTimeout creates a new PutPolicyPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPolicyPathParamsWithTimeout(timeout time.Duration) *PutPolicyPathParams {
	var ()
	return &PutPolicyPathParams{

		timeout: timeout,
	}
}

// NewPutPolicyPathParamsWithContext creates a new PutPolicyPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPolicyPathParamsWithContext(ctx context.Context) *PutPolicyPathParams {
	var ()
	return &PutPolicyPathParams{

		Context: ctx,
	}
}

// NewPutPolicyPathParamsWithHTTPClient creates a new PutPolicyPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPolicyPathParamsWithHTTPClient(client *http.Client) *PutPolicyPathParams {
	var ()
	return &PutPolicyPathParams{
		HTTPClient: client,
	}
}

/*PutPolicyPathParams contains all the parameters to send to the API endpoint
for the put policy path operation typically these are written to a http.Request
*/
type PutPolicyPathParams struct {

	/*Path
	  Path to policy node

	*/
	Path string
	/*Policy
	  Policy tree or subtree

	*/
	Policy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put policy path params
func (o *PutPolicyPathParams) WithTimeout(timeout time.Duration) *PutPolicyPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put policy path params
func (o *PutPolicyPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put policy path params
func (o *PutPolicyPathParams) WithContext(ctx context.Context) *PutPolicyPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put policy path params
func (o *PutPolicyPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put policy path params
func (o *PutPolicyPathParams) WithHTTPClient(client *http.Client) *PutPolicyPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put policy path params
func (o *PutPolicyPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the put policy path params
func (o *PutPolicyPathParams) WithPath(path string) *PutPolicyPathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the put policy path params
func (o *PutPolicyPathParams) SetPath(path string) {
	o.Path = path
}

// WithPolicy adds the policy to the put policy path params
func (o *PutPolicyPathParams) WithPolicy(policy *string) *PutPolicyPathParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the put policy path params
func (o *PutPolicyPathParams) SetPolicy(policy *string) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PutPolicyPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Policy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
