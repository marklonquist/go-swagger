// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PetUploadImageHandlerFunc turns a function with the right signature into a pet upload image handler
type PetUploadImageHandlerFunc func(PetUploadImageParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PetUploadImageHandlerFunc) Handle(params PetUploadImageParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PetUploadImageHandler interface for that can handle valid pet upload image params
type PetUploadImageHandler interface {
	Handle(PetUploadImageParams, interface{}) middleware.Responder
}

// NewPetUploadImage creates a new http.Handler for the pet upload image operation
func NewPetUploadImage(ctx *middleware.Context, handler PetUploadImageHandler) *PetUploadImage {
	return &PetUploadImage{Context: ctx, Handler: handler}
}

/*PetUploadImage swagger:route POST /pet/{petId}/image pet petUploadImage

uploads an image

*/
type PetUploadImage struct {
	Context *middleware.Context
	Handler PetUploadImageHandler
}

func (o *PetUploadImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPetUploadImageParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
