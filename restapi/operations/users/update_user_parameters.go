package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateUserParams creates a new UpdateUserParams object
// with the default values initialized.
func NewUpdateUserParams() UpdateUserParams {
	var ()
	return UpdateUserParams{}
}

// UpdateUserParams contains all the bound params for the update user operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateUser
type UpdateUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*User ID
	  Required: true
	  In: path
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("user_id", "path", "int64", raw)
	}
	o.UserID = value

	return nil
}
