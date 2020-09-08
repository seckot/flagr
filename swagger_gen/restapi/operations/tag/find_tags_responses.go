// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/checkr/flagr/swagger_gen/models"
)

// FindTagsOKCode is the HTTP code returned for type FindTagsOK
const FindTagsOKCode int = 200

/*FindTagsOK tag ordered by tagID

swagger:response findTagsOK
*/
type FindTagsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tag `json:"body,omitempty"`
}

// NewFindTagsOK creates FindTagsOK with default headers values
func NewFindTagsOK() *FindTagsOK {

	return &FindTagsOK{}
}

// WithPayload adds the payload to the find tags o k response
func (o *FindTagsOK) WithPayload(payload []*models.Tag) *FindTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find tags o k response
func (o *FindTagsOK) SetPayload(payload []*models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Tag, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*FindTagsDefault generic error response

swagger:response findTagsDefault
*/
type FindTagsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindTagsDefault creates FindTagsDefault with default headers values
func NewFindTagsDefault(code int) *FindTagsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindTagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find tags default response
func (o *FindTagsDefault) WithStatusCode(code int) *FindTagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find tags default response
func (o *FindTagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find tags default response
func (o *FindTagsDefault) WithPayload(payload *models.Error) *FindTagsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find tags default response
func (o *FindTagsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}