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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasReportsListReader is a Reader for the ExtrasReportsList structure.
type ExtrasReportsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasReportsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasReportsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasReportsListOK creates a ExtrasReportsListOK with default headers values
func NewExtrasReportsListOK() *ExtrasReportsListOK {
	return &ExtrasReportsListOK{}
}

/*
ExtrasReportsListOK describes a response with status code 200, with default header values.

ExtrasReportsListOK extras reports list o k
*/
type ExtrasReportsListOK struct {
}

// IsSuccess returns true when this extras reports list o k response has a 2xx status code
func (o *ExtrasReportsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras reports list o k response has a 3xx status code
func (o *ExtrasReportsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras reports list o k response has a 4xx status code
func (o *ExtrasReportsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras reports list o k response has a 5xx status code
func (o *ExtrasReportsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras reports list o k response a status code equal to that given
func (o *ExtrasReportsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasReportsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/reports/][%d] extrasReportsListOK ", 200)
}

func (o *ExtrasReportsListOK) String() string {
	return fmt.Sprintf("[GET /extras/reports/][%d] extrasReportsListOK ", 200)
}

func (o *ExtrasReportsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
