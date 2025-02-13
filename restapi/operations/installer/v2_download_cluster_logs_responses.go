// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2DownloadClusterLogsOKCode is the HTTP code returned for type V2DownloadClusterLogsOK
const V2DownloadClusterLogsOKCode int = 200

/*V2DownloadClusterLogsOK Success.

swagger:response v2DownloadClusterLogsOK
*/
type V2DownloadClusterLogsOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsOK creates V2DownloadClusterLogsOK with default headers values
func NewV2DownloadClusterLogsOK() *V2DownloadClusterLogsOK {

	return &V2DownloadClusterLogsOK{}
}

// WithPayload adds the payload to the v2 download cluster logs o k response
func (o *V2DownloadClusterLogsOK) WithPayload(payload io.ReadCloser) *V2DownloadClusterLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs o k response
func (o *V2DownloadClusterLogsOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2DownloadClusterLogsUnauthorizedCode is the HTTP code returned for type V2DownloadClusterLogsUnauthorized
const V2DownloadClusterLogsUnauthorizedCode int = 401

/*V2DownloadClusterLogsUnauthorized Unauthorized.

swagger:response v2DownloadClusterLogsUnauthorized
*/
type V2DownloadClusterLogsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsUnauthorized creates V2DownloadClusterLogsUnauthorized with default headers values
func NewV2DownloadClusterLogsUnauthorized() *V2DownloadClusterLogsUnauthorized {

	return &V2DownloadClusterLogsUnauthorized{}
}

// WithPayload adds the payload to the v2 download cluster logs unauthorized response
func (o *V2DownloadClusterLogsUnauthorized) WithPayload(payload *models.InfraError) *V2DownloadClusterLogsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs unauthorized response
func (o *V2DownloadClusterLogsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterLogsForbiddenCode is the HTTP code returned for type V2DownloadClusterLogsForbidden
const V2DownloadClusterLogsForbiddenCode int = 403

/*V2DownloadClusterLogsForbidden Forbidden.

swagger:response v2DownloadClusterLogsForbidden
*/
type V2DownloadClusterLogsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsForbidden creates V2DownloadClusterLogsForbidden with default headers values
func NewV2DownloadClusterLogsForbidden() *V2DownloadClusterLogsForbidden {

	return &V2DownloadClusterLogsForbidden{}
}

// WithPayload adds the payload to the v2 download cluster logs forbidden response
func (o *V2DownloadClusterLogsForbidden) WithPayload(payload *models.InfraError) *V2DownloadClusterLogsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs forbidden response
func (o *V2DownloadClusterLogsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterLogsNotFoundCode is the HTTP code returned for type V2DownloadClusterLogsNotFound
const V2DownloadClusterLogsNotFoundCode int = 404

/*V2DownloadClusterLogsNotFound Error.

swagger:response v2DownloadClusterLogsNotFound
*/
type V2DownloadClusterLogsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsNotFound creates V2DownloadClusterLogsNotFound with default headers values
func NewV2DownloadClusterLogsNotFound() *V2DownloadClusterLogsNotFound {

	return &V2DownloadClusterLogsNotFound{}
}

// WithPayload adds the payload to the v2 download cluster logs not found response
func (o *V2DownloadClusterLogsNotFound) WithPayload(payload *models.Error) *V2DownloadClusterLogsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs not found response
func (o *V2DownloadClusterLogsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterLogsMethodNotAllowedCode is the HTTP code returned for type V2DownloadClusterLogsMethodNotAllowed
const V2DownloadClusterLogsMethodNotAllowedCode int = 405

/*V2DownloadClusterLogsMethodNotAllowed Method Not Allowed.

swagger:response v2DownloadClusterLogsMethodNotAllowed
*/
type V2DownloadClusterLogsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsMethodNotAllowed creates V2DownloadClusterLogsMethodNotAllowed with default headers values
func NewV2DownloadClusterLogsMethodNotAllowed() *V2DownloadClusterLogsMethodNotAllowed {

	return &V2DownloadClusterLogsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 download cluster logs method not allowed response
func (o *V2DownloadClusterLogsMethodNotAllowed) WithPayload(payload *models.Error) *V2DownloadClusterLogsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs method not allowed response
func (o *V2DownloadClusterLogsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterLogsConflictCode is the HTTP code returned for type V2DownloadClusterLogsConflict
const V2DownloadClusterLogsConflictCode int = 409

/*V2DownloadClusterLogsConflict Error.

swagger:response v2DownloadClusterLogsConflict
*/
type V2DownloadClusterLogsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsConflict creates V2DownloadClusterLogsConflict with default headers values
func NewV2DownloadClusterLogsConflict() *V2DownloadClusterLogsConflict {

	return &V2DownloadClusterLogsConflict{}
}

// WithPayload adds the payload to the v2 download cluster logs conflict response
func (o *V2DownloadClusterLogsConflict) WithPayload(payload *models.Error) *V2DownloadClusterLogsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs conflict response
func (o *V2DownloadClusterLogsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterLogsInternalServerErrorCode is the HTTP code returned for type V2DownloadClusterLogsInternalServerError
const V2DownloadClusterLogsInternalServerErrorCode int = 500

/*V2DownloadClusterLogsInternalServerError Error.

swagger:response v2DownloadClusterLogsInternalServerError
*/
type V2DownloadClusterLogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterLogsInternalServerError creates V2DownloadClusterLogsInternalServerError with default headers values
func NewV2DownloadClusterLogsInternalServerError() *V2DownloadClusterLogsInternalServerError {

	return &V2DownloadClusterLogsInternalServerError{}
}

// WithPayload adds the payload to the v2 download cluster logs internal server error response
func (o *V2DownloadClusterLogsInternalServerError) WithPayload(payload *models.Error) *V2DownloadClusterLogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster logs internal server error response
func (o *V2DownloadClusterLogsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
