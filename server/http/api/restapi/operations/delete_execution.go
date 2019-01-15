// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteExecutionHandlerFunc turns a function with the right signature into a delete execution handler
type DeleteExecutionHandlerFunc func(DeleteExecutionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteExecutionHandlerFunc) Handle(params DeleteExecutionParams) middleware.Responder {
	return fn(params)
}

// DeleteExecutionHandler interface for that can handle valid delete execution params
type DeleteExecutionHandler interface {
	Handle(DeleteExecutionParams) middleware.Responder
}

// NewDeleteExecution creates a new http.Handler for the delete execution operation
func NewDeleteExecution(ctx *middleware.Context, handler DeleteExecutionHandler) *DeleteExecution {
	return &DeleteExecution{Context: ctx, Handler: handler}
}

/*DeleteExecution swagger:route DELETE /projects/{projectID}/exec/{jobID} deleteExecution

Cancel Execution

*/
type DeleteExecution struct {
	Context *middleware.Context
	Handler DeleteExecutionHandler
}

func (o *DeleteExecution) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteExecutionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
