package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ShareRequest share request
// swagger:model ShareRequest
type ShareRequest struct {

	// allow details
	AllowDetails *bool `json:"allowDetails,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// share proto
	ShareProto string `json:"shareProto,omitempty"`

	// share type
	ShareType string `json:"shareType,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`
}

// Validate validates this share request
func (m *ShareRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
