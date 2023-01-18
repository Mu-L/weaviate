//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsPropertiesAddOKCode is the HTTP code returned for type SchemaObjectsPropertiesAddOK
const SchemaObjectsPropertiesAddOKCode int = 200

/*
SchemaObjectsPropertiesAddOK Added the property.

swagger:response schemaObjectsPropertiesAddOK
*/
type SchemaObjectsPropertiesAddOK struct {

	/*
	  In: Body
	*/
	Payload *models.Property `json:"body,omitempty"`
}

// NewSchemaObjectsPropertiesAddOK creates SchemaObjectsPropertiesAddOK with default headers values
func NewSchemaObjectsPropertiesAddOK() *SchemaObjectsPropertiesAddOK {

	return &SchemaObjectsPropertiesAddOK{}
}

// WithPayload adds the payload to the schema objects properties add o k response
func (o *SchemaObjectsPropertiesAddOK) WithPayload(payload *models.Property) *SchemaObjectsPropertiesAddOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects properties add o k response
func (o *SchemaObjectsPropertiesAddOK) SetPayload(payload *models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsPropertiesAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsPropertiesAddUnauthorizedCode is the HTTP code returned for type SchemaObjectsPropertiesAddUnauthorized
const SchemaObjectsPropertiesAddUnauthorizedCode int = 401

/*
SchemaObjectsPropertiesAddUnauthorized Unauthorized or invalid credentials.

swagger:response schemaObjectsPropertiesAddUnauthorized
*/
type SchemaObjectsPropertiesAddUnauthorized struct {
}

// NewSchemaObjectsPropertiesAddUnauthorized creates SchemaObjectsPropertiesAddUnauthorized with default headers values
func NewSchemaObjectsPropertiesAddUnauthorized() *SchemaObjectsPropertiesAddUnauthorized {

	return &SchemaObjectsPropertiesAddUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaObjectsPropertiesAddUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaObjectsPropertiesAddForbiddenCode is the HTTP code returned for type SchemaObjectsPropertiesAddForbidden
const SchemaObjectsPropertiesAddForbiddenCode int = 403

/*
SchemaObjectsPropertiesAddForbidden Forbidden

swagger:response schemaObjectsPropertiesAddForbidden
*/
type SchemaObjectsPropertiesAddForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsPropertiesAddForbidden creates SchemaObjectsPropertiesAddForbidden with default headers values
func NewSchemaObjectsPropertiesAddForbidden() *SchemaObjectsPropertiesAddForbidden {

	return &SchemaObjectsPropertiesAddForbidden{}
}

// WithPayload adds the payload to the schema objects properties add forbidden response
func (o *SchemaObjectsPropertiesAddForbidden) WithPayload(payload *models.ErrorResponse) *SchemaObjectsPropertiesAddForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects properties add forbidden response
func (o *SchemaObjectsPropertiesAddForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsPropertiesAddForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsPropertiesAddUnprocessableEntityCode is the HTTP code returned for type SchemaObjectsPropertiesAddUnprocessableEntity
const SchemaObjectsPropertiesAddUnprocessableEntityCode int = 422

/*
SchemaObjectsPropertiesAddUnprocessableEntity Invalid property.

swagger:response schemaObjectsPropertiesAddUnprocessableEntity
*/
type SchemaObjectsPropertiesAddUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsPropertiesAddUnprocessableEntity creates SchemaObjectsPropertiesAddUnprocessableEntity with default headers values
func NewSchemaObjectsPropertiesAddUnprocessableEntity() *SchemaObjectsPropertiesAddUnprocessableEntity {

	return &SchemaObjectsPropertiesAddUnprocessableEntity{}
}

// WithPayload adds the payload to the schema objects properties add unprocessable entity response
func (o *SchemaObjectsPropertiesAddUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaObjectsPropertiesAddUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects properties add unprocessable entity response
func (o *SchemaObjectsPropertiesAddUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsPropertiesAddUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsPropertiesAddInternalServerErrorCode is the HTTP code returned for type SchemaObjectsPropertiesAddInternalServerError
const SchemaObjectsPropertiesAddInternalServerErrorCode int = 500

/*
SchemaObjectsPropertiesAddInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaObjectsPropertiesAddInternalServerError
*/
type SchemaObjectsPropertiesAddInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsPropertiesAddInternalServerError creates SchemaObjectsPropertiesAddInternalServerError with default headers values
func NewSchemaObjectsPropertiesAddInternalServerError() *SchemaObjectsPropertiesAddInternalServerError {

	return &SchemaObjectsPropertiesAddInternalServerError{}
}

// WithPayload adds the payload to the schema objects properties add internal server error response
func (o *SchemaObjectsPropertiesAddInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaObjectsPropertiesAddInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects properties add internal server error response
func (o *SchemaObjectsPropertiesAddInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsPropertiesAddInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
