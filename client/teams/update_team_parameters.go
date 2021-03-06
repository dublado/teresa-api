package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// NewUpdateTeamParams creates a new UpdateTeamParams object
// with the default values initialized.
func NewUpdateTeamParams() *UpdateTeamParams {
	var ()
	return &UpdateTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeamParamsWithTimeout creates a new UpdateTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTeamParamsWithTimeout(timeout time.Duration) *UpdateTeamParams {
	var ()
	return &UpdateTeamParams{

		timeout: timeout,
	}
}

/*UpdateTeamParams contains all the parameters to send to the API endpoint
for the update team operation typically these are written to a http.Request
*/
type UpdateTeamParams struct {

	/*Body*/
	Body *models.Team
	/*TeamID
	  Team ID

	*/
	TeamID int64

	timeout time.Duration
}

// WithBody adds the body to the update team params
func (o *UpdateTeamParams) WithBody(body *models.Team) *UpdateTeamParams {
	o.Body = body
	return o
}

// WithTeamID adds the teamID to the update team params
func (o *UpdateTeamParams) WithTeamID(teamID int64) *UpdateTeamParams {
	o.TeamID = teamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.Team)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
