// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewListClusterBackupStorageLocationBucketObjectsParams creates a new ListClusterBackupStorageLocationBucketObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClusterBackupStorageLocationBucketObjectsParams() *ListClusterBackupStorageLocationBucketObjectsParams {
	return &ListClusterBackupStorageLocationBucketObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterBackupStorageLocationBucketObjectsParamsWithTimeout creates a new ListClusterBackupStorageLocationBucketObjectsParams object
// with the ability to set a timeout on a request.
func NewListClusterBackupStorageLocationBucketObjectsParamsWithTimeout(timeout time.Duration) *ListClusterBackupStorageLocationBucketObjectsParams {
	return &ListClusterBackupStorageLocationBucketObjectsParams{
		timeout: timeout,
	}
}

// NewListClusterBackupStorageLocationBucketObjectsParamsWithContext creates a new ListClusterBackupStorageLocationBucketObjectsParams object
// with the ability to set a context for a request.
func NewListClusterBackupStorageLocationBucketObjectsParamsWithContext(ctx context.Context) *ListClusterBackupStorageLocationBucketObjectsParams {
	return &ListClusterBackupStorageLocationBucketObjectsParams{
		Context: ctx,
	}
}

// NewListClusterBackupStorageLocationBucketObjectsParamsWithHTTPClient creates a new ListClusterBackupStorageLocationBucketObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClusterBackupStorageLocationBucketObjectsParamsWithHTTPClient(client *http.Client) *ListClusterBackupStorageLocationBucketObjectsParams {
	return &ListClusterBackupStorageLocationBucketObjectsParams{
		HTTPClient: client,
	}
}

/*
ListClusterBackupStorageLocationBucketObjectsParams contains all the parameters to send to the API endpoint

	for the list cluster backup storage location bucket objects operation.

	Typically these are written to a http.Request.
*/
type ListClusterBackupStorageLocationBucketObjectsParams struct {

	// CbslName.
	ClusterBackupStorageLocationName string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list cluster backup storage location bucket objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithDefaults() *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list cluster backup storage location bucket objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithTimeout(timeout time.Duration) *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithContext(ctx context.Context) *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithHTTPClient(client *http.Client) *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterBackupStorageLocationName adds the cbslName to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithClusterBackupStorageLocationName(cbslName string) *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetClusterBackupStorageLocationName(cbslName)
	return o
}

// SetClusterBackupStorageLocationName adds the cbslName to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetClusterBackupStorageLocationName(cbslName string) {
	o.ClusterBackupStorageLocationName = cbslName
}

// WithProjectID adds the projectID to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WithProjectID(projectID string) *ListClusterBackupStorageLocationBucketObjectsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list cluster backup storage location bucket objects params
func (o *ListClusterBackupStorageLocationBucketObjectsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterBackupStorageLocationBucketObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cbsl_name
	if err := r.SetPathParam("cbsl_name", o.ClusterBackupStorageLocationName); err != nil {
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
