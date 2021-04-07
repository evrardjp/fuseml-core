// Code generated by go-swagger; DO NOT EDIT.

package runnable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterRunnableHandlerFunc turns a function with the right signature into a register runnable handler
type RegisterRunnableHandlerFunc func(RegisterRunnableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterRunnableHandlerFunc) Handle(params RegisterRunnableParams) middleware.Responder {
	return fn(params)
}

// RegisterRunnableHandler interface for that can handle valid register runnable params
type RegisterRunnableHandler interface {
	Handle(RegisterRunnableParams) middleware.Responder
}

// NewRegisterRunnable creates a new http.Handler for the register runnable operation
func NewRegisterRunnable(ctx *middleware.Context, handler RegisterRunnableHandler) *RegisterRunnable {
	return &RegisterRunnable{Context: ctx, Handler: handler}
}

/* RegisterRunnable swagger:route POST /runnables runnable registerRunnable

Register runnable

Register a runnable with the FuseML runnable store

*/
type RegisterRunnable struct {
	Context *middleware.Context
	Handler RegisterRunnableHandler
}

func (o *RegisterRunnable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterRunnableParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
