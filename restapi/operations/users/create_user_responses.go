package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*CreateUserCreated Newly created user

swagger:response createUserCreated
*/
type CreateUserCreated struct {

	// In: body
	Payload *models.User `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *models.User) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUserBadRequest User sent a bad request

swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {

	// In: body
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

// WithPayload adds the payload to the create user bad request response
func (o *CreateUserBadRequest) WithPayload(payload *models.BadRequest) *CreateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user bad request response
func (o *CreateUserBadRequest) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUserUnauthorized User not authorized

swagger:response createUserUnauthorized
*/
type CreateUserUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewCreateUserUnauthorized creates CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

// WithPayload adds the payload to the create user unauthorized response
func (o *CreateUserUnauthorized) WithPayload(payload *models.Unauthorized) *CreateUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user unauthorized response
func (o *CreateUserUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUserForbidden User does not have the credentials to access this resource


swagger:response createUserForbidden
*/
type CreateUserForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewCreateUserForbidden creates CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

// WithPayload adds the payload to the create user forbidden response
func (o *CreateUserForbidden) WithPayload(payload *models.Unauthorized) *CreateUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user forbidden response
func (o *CreateUserForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUserDefault Error

swagger:response createUserDefault
*/
type CreateUserDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUserDefault creates CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create user default response
func (o *CreateUserDefault) WithStatusCode(code int) *CreateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create user default response
func (o *CreateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create user default response
func (o *CreateUserDefault) WithPayload(payload *models.Error) *CreateUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user default response
func (o *CreateUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
