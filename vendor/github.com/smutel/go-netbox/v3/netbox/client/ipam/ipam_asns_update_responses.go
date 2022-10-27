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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// IpamAsnsUpdateReader is a Reader for the IpamAsnsUpdate structure.
type IpamAsnsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAsnsUpdateOK creates a IpamAsnsUpdateOK with default headers values
func NewIpamAsnsUpdateOK() *IpamAsnsUpdateOK {
	return &IpamAsnsUpdateOK{}
}

/*
IpamAsnsUpdateOK describes a response with status code 200, with default header values.

IpamAsnsUpdateOK ipam asns update o k
*/
type IpamAsnsUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns update o k response has a 2xx status code
func (o *IpamAsnsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns update o k response has a 3xx status code
func (o *IpamAsnsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns update o k response has a 4xx status code
func (o *IpamAsnsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns update o k response has a 5xx status code
func (o *IpamAsnsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns update o k response a status code equal to that given
func (o *IpamAsnsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamAsnsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipamAsnsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipamAsnsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
