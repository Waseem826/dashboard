// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListVSphereVMGroupsParams creates a new ListVSphereVMGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVSphereVMGroupsParams() *ListVSphereVMGroupsParams {
	return &ListVSphereVMGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVSphereVMGroupsParamsWithTimeout creates a new ListVSphereVMGroupsParams object
// with the ability to set a timeout on a request.
func NewListVSphereVMGroupsParamsWithTimeout(timeout time.Duration) *ListVSphereVMGroupsParams {
	return &ListVSphereVMGroupsParams{
		timeout: timeout,
	}
}

// NewListVSphereVMGroupsParamsWithContext creates a new ListVSphereVMGroupsParams object
// with the ability to set a context for a request.
func NewListVSphereVMGroupsParamsWithContext(ctx context.Context) *ListVSphereVMGroupsParams {
	return &ListVSphereVMGroupsParams{
		Context: ctx,
	}
}

// NewListVSphereVMGroupsParamsWithHTTPClient creates a new ListVSphereVMGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVSphereVMGroupsParamsWithHTTPClient(client *http.Client) *ListVSphereVMGroupsParams {
	return &ListVSphereVMGroupsParams{
		HTTPClient: client,
	}
}

/*
ListVSphereVMGroupsParams contains all the parameters to send to the API endpoint

	for the list v sphere VM groups operation.

	Typically these are written to a http.Request.
*/
type ListVSphereVMGroupsParams struct {

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Password.
	Password *string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list v sphere VM groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereVMGroupsParams) WithDefaults() *ListVSphereVMGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list v sphere VM groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereVMGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithTimeout(timeout time.Duration) *ListVSphereVMGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithContext(ctx context.Context) *ListVSphereVMGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithHTTPClient(client *http.Client) *ListVSphereVMGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithCredential(credential *string) *ListVSphereVMGroupsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithDatacenterName(datacenterName *string) *ListVSphereVMGroupsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithPassword adds the password to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithPassword(password *string) *ListVSphereVMGroupsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) WithUsername(username *string) *ListVSphereVMGroupsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list v sphere VM groups params
func (o *ListVSphereVMGroupsParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListVSphereVMGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
