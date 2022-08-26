// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomLink custom link
//
// swagger:model CustomLink
type CustomLink struct {

	// Button class
	//
	// The class of the first link in a group will be used for the dropdown button
	// Enum: [outline-dark blue indigo purple pink red orange yellow green teal cyan gray black white ghost-dark]
	ButtonClass string `json:"button_class,omitempty"`

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// Group name
	//
	// Links with the same group will appear as a dropdown menu
	// Max Length: 50
	GroupName string `json:"group_name,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Link text
	//
	// Jinja2 template code for link text
	// Required: true
	// Max Length: 500
	// Min Length: 1
	LinkText *string `json:"link_text"`

	// Link URL
	//
	// Jinja2 template code for link URL
	// Required: true
	// Max Length: 500
	// Min Length: 1
	LinkURL *string `json:"link_url"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// New window
	//
	// Force link to open in a new window
	NewWindow bool `json:"new_window,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this custom link
func (m *CustomLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateButtonClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customLinkTypeButtonClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["outline-dark","blue","indigo","purple","pink","red","orange","yellow","green","teal","cyan","gray","black","white","ghost-dark"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customLinkTypeButtonClassPropEnum = append(customLinkTypeButtonClassPropEnum, v)
	}
}

const (

	// CustomLinkButtonClassOutlineDashDark captures enum value "outline-dark"
	CustomLinkButtonClassOutlineDashDark string = "outline-dark"

	// CustomLinkButtonClassBlue captures enum value "blue"
	CustomLinkButtonClassBlue string = "blue"

	// CustomLinkButtonClassIndigo captures enum value "indigo"
	CustomLinkButtonClassIndigo string = "indigo"

	// CustomLinkButtonClassPurple captures enum value "purple"
	CustomLinkButtonClassPurple string = "purple"

	// CustomLinkButtonClassPink captures enum value "pink"
	CustomLinkButtonClassPink string = "pink"

	// CustomLinkButtonClassRed captures enum value "red"
	CustomLinkButtonClassRed string = "red"

	// CustomLinkButtonClassOrange captures enum value "orange"
	CustomLinkButtonClassOrange string = "orange"

	// CustomLinkButtonClassYellow captures enum value "yellow"
	CustomLinkButtonClassYellow string = "yellow"

	// CustomLinkButtonClassGreen captures enum value "green"
	CustomLinkButtonClassGreen string = "green"

	// CustomLinkButtonClassTeal captures enum value "teal"
	CustomLinkButtonClassTeal string = "teal"

	// CustomLinkButtonClassCyan captures enum value "cyan"
	CustomLinkButtonClassCyan string = "cyan"

	// CustomLinkButtonClassGray captures enum value "gray"
	CustomLinkButtonClassGray string = "gray"

	// CustomLinkButtonClassBlack captures enum value "black"
	CustomLinkButtonClassBlack string = "black"

	// CustomLinkButtonClassWhite captures enum value "white"
	CustomLinkButtonClassWhite string = "white"

	// CustomLinkButtonClassGhostDashDark captures enum value "ghost-dark"
	CustomLinkButtonClassGhostDashDark string = "ghost-dark"
)

// prop value enum
func (m *CustomLink) validateButtonClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customLinkTypeButtonClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomLink) validateButtonClass(formats strfmt.Registry) error {
	if swag.IsZero(m.ButtonClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateButtonClassEnum("button_class", "body", m.ButtonClass); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateGroupName(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupName) { // not required
		return nil
	}

	if err := validate.MaxLength("group_name", "body", m.GroupName, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateLinkText(formats strfmt.Registry) error {

	if err := validate.Required("link_text", "body", m.LinkText); err != nil {
		return err
	}

	if err := validate.MinLength("link_text", "body", *m.LinkText, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("link_text", "body", *m.LinkText, 500); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateLinkURL(formats strfmt.Registry) error {

	if err := validate.Required("link_url", "body", m.LinkURL); err != nil {
		return err
	}

	if err := validate.MinLength("link_url", "body", *m.LinkURL, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("link_url", "body", *m.LinkURL, 500); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this custom link based on the context it is used
func (m *CustomLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomLink) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomLink) UnmarshalBinary(b []byte) error {
	var res CustomLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
