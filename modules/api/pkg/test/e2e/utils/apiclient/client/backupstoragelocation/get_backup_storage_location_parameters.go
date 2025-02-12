// Code generated by go-swagger; DO NOT EDIT.

package backupstoragelocation

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

// NewGetBackupStorageLocationParams creates a new GetBackupStorageLocationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupStorageLocationParams() *GetBackupStorageLocationParams {
	return &GetBackupStorageLocationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupStorageLocationParamsWithTimeout creates a new GetBackupStorageLocationParams object
// with the ability to set a timeout on a request.
func NewGetBackupStorageLocationParamsWithTimeout(timeout time.Duration) *GetBackupStorageLocationParams {
	return &GetBackupStorageLocationParams{
		timeout: timeout,
	}
}

// NewGetBackupStorageLocationParamsWithContext creates a new GetBackupStorageLocationParams object
// with the ability to set a context for a request.
func NewGetBackupStorageLocationParamsWithContext(ctx context.Context) *GetBackupStorageLocationParams {
	return &GetBackupStorageLocationParams{
		Context: ctx,
	}
}

// NewGetBackupStorageLocationParamsWithHTTPClient creates a new GetBackupStorageLocationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupStorageLocationParamsWithHTTPClient(client *http.Client) *GetBackupStorageLocationParams {
	return &GetBackupStorageLocationParams{
		HTTPClient: client,
	}
}

/*
GetBackupStorageLocationParams contains all the parameters to send to the API endpoint

	for the get backup storage location operation.

	Typically these are written to a http.Request.
*/
type GetBackupStorageLocationParams struct {

	// BslName.
	BSLName string

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup storage location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupStorageLocationParams) WithDefaults() *GetBackupStorageLocationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup storage location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupStorageLocationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithTimeout(timeout time.Duration) *GetBackupStorageLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithContext(ctx context.Context) *GetBackupStorageLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithHTTPClient(client *http.Client) *GetBackupStorageLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBSLName adds the bslName to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithBSLName(bslName string) *GetBackupStorageLocationParams {
	o.SetBSLName(bslName)
	return o
}

// SetBSLName adds the bslName to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetBSLName(bslName string) {
	o.BSLName = bslName
}

// WithClusterID adds the clusterID to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithClusterID(clusterID string) *GetBackupStorageLocationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get backup storage location params
func (o *GetBackupStorageLocationParams) WithProjectID(projectID string) *GetBackupStorageLocationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get backup storage location params
func (o *GetBackupStorageLocationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupStorageLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bsl_name
	if err := r.SetPathParam("bsl_name", o.BSLName); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
