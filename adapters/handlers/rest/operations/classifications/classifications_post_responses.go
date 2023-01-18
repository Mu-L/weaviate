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

package classifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ClassificationsPostCreatedCode is the HTTP code returned for type ClassificationsPostCreated
const ClassificationsPostCreatedCode int = 201

/*
ClassificationsPostCreated Successfully started classification.

swagger:response classificationsPostCreated
*/
type ClassificationsPostCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Classification `json:"body,omitempty"`
}

// NewClassificationsPostCreated creates ClassificationsPostCreated with default headers values
func NewClassificationsPostCreated() *ClassificationsPostCreated {

	return &ClassificationsPostCreated{}
}

// WithPayload adds the payload to the classifications post created response
func (o *ClassificationsPostCreated) WithPayload(payload *models.Classification) *ClassificationsPostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications post created response
func (o *ClassificationsPostCreated) SetPayload(payload *models.Classification) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsPostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClassificationsPostBadRequestCode is the HTTP code returned for type ClassificationsPostBadRequest
const ClassificationsPostBadRequestCode int = 400

/*
ClassificationsPostBadRequest Incorrect request

swagger:response classificationsPostBadRequest
*/
type ClassificationsPostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewClassificationsPostBadRequest creates ClassificationsPostBadRequest with default headers values
func NewClassificationsPostBadRequest() *ClassificationsPostBadRequest {

	return &ClassificationsPostBadRequest{}
}

// WithPayload adds the payload to the classifications post bad request response
func (o *ClassificationsPostBadRequest) WithPayload(payload *models.ErrorResponse) *ClassificationsPostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications post bad request response
func (o *ClassificationsPostBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsPostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClassificationsPostUnauthorizedCode is the HTTP code returned for type ClassificationsPostUnauthorized
const ClassificationsPostUnauthorizedCode int = 401

/*
ClassificationsPostUnauthorized Unauthorized or invalid credentials.

swagger:response classificationsPostUnauthorized
*/
type ClassificationsPostUnauthorized struct {
}

// NewClassificationsPostUnauthorized creates ClassificationsPostUnauthorized with default headers values
func NewClassificationsPostUnauthorized() *ClassificationsPostUnauthorized {

	return &ClassificationsPostUnauthorized{}
}

// WriteResponse to the client
func (o *ClassificationsPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ClassificationsPostForbiddenCode is the HTTP code returned for type ClassificationsPostForbidden
const ClassificationsPostForbiddenCode int = 403

/*
ClassificationsPostForbidden Forbidden

swagger:response classificationsPostForbidden
*/
type ClassificationsPostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewClassificationsPostForbidden creates ClassificationsPostForbidden with default headers values
func NewClassificationsPostForbidden() *ClassificationsPostForbidden {

	return &ClassificationsPostForbidden{}
}

// WithPayload adds the payload to the classifications post forbidden response
func (o *ClassificationsPostForbidden) WithPayload(payload *models.ErrorResponse) *ClassificationsPostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications post forbidden response
func (o *ClassificationsPostForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClassificationsPostInternalServerErrorCode is the HTTP code returned for type ClassificationsPostInternalServerError
const ClassificationsPostInternalServerErrorCode int = 500

/*
ClassificationsPostInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response classificationsPostInternalServerError
*/
type ClassificationsPostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewClassificationsPostInternalServerError creates ClassificationsPostInternalServerError with default headers values
func NewClassificationsPostInternalServerError() *ClassificationsPostInternalServerError {

	return &ClassificationsPostInternalServerError{}
}

// WithPayload adds the payload to the classifications post internal server error response
func (o *ClassificationsPostInternalServerError) WithPayload(payload *models.ErrorResponse) *ClassificationsPostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications post internal server error response
func (o *ClassificationsPostInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsPostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
