// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Alibaba alibaba
//
// swagger:model Alibaba
type Alibaba struct {

	// The Access Key ID used to authenticate against Alibaba.
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// The Access Key Secret used to authenticate against Alibaba.
	AccessKeySecret string `json:"accessKeySecret,omitempty"`

	// If datacenter is set, this preset is only applicable to the
	// configured datacenter.
	Datacenter string `json:"datacenter,omitempty"`

	// Only enabled presets will be available in the KKP dashboard.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this alibaba
func (m *Alibaba) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alibaba based on context it is used
func (m *Alibaba) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Alibaba) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alibaba) UnmarshalBinary(b []byte) error {
	var res Alibaba
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
