// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2GetClusterReader is a Reader for the V2GetCluster structure.
type V2GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2GetClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2GetClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2GetClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2GetClusterMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2GetClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2GetClusterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2GetClusterOK creates a V2GetClusterOK with default headers values
func NewV2GetClusterOK() *V2GetClusterOK {
	return &V2GetClusterOK{}
}

/*V2GetClusterOK handles this case with default header values.

Success.
*/
type V2GetClusterOK struct {
	Payload *models.Cluster
}

func (o *V2GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterOK  %+v", 200, o.Payload)
}

func (o *V2GetClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *V2GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterUnauthorized creates a V2GetClusterUnauthorized with default headers values
func NewV2GetClusterUnauthorized() *V2GetClusterUnauthorized {
	return &V2GetClusterUnauthorized{}
}

/*V2GetClusterUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2GetClusterUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2GetClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *V2GetClusterUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterForbidden creates a V2GetClusterForbidden with default headers values
func NewV2GetClusterForbidden() *V2GetClusterForbidden {
	return &V2GetClusterForbidden{}
}

/*V2GetClusterForbidden handles this case with default header values.

Forbidden.
*/
type V2GetClusterForbidden struct {
	Payload *models.InfraError
}

func (o *V2GetClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterForbidden  %+v", 403, o.Payload)
}

func (o *V2GetClusterForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterNotFound creates a V2GetClusterNotFound with default headers values
func NewV2GetClusterNotFound() *V2GetClusterNotFound {
	return &V2GetClusterNotFound{}
}

/*V2GetClusterNotFound handles this case with default header values.

Error.
*/
type V2GetClusterNotFound struct {
	Payload *models.Error
}

func (o *V2GetClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterNotFound  %+v", 404, o.Payload)
}

func (o *V2GetClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterMethodNotAllowed creates a V2GetClusterMethodNotAllowed with default headers values
func NewV2GetClusterMethodNotAllowed() *V2GetClusterMethodNotAllowed {
	return &V2GetClusterMethodNotAllowed{}
}

/*V2GetClusterMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2GetClusterMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2GetClusterMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2GetClusterMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetClusterMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterInternalServerError creates a V2GetClusterInternalServerError with default headers values
func NewV2GetClusterInternalServerError() *V2GetClusterInternalServerError {
	return &V2GetClusterInternalServerError{}
}

/*V2GetClusterInternalServerError handles this case with default header values.

Error.
*/
type V2GetClusterInternalServerError struct {
	Payload *models.Error
}

func (o *V2GetClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetClusterServiceUnavailable creates a V2GetClusterServiceUnavailable with default headers values
func NewV2GetClusterServiceUnavailable() *V2GetClusterServiceUnavailable {
	return &V2GetClusterServiceUnavailable{}
}

/*V2GetClusterServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2GetClusterServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2GetClusterServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}][%d] v2GetClusterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2GetClusterServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetClusterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
