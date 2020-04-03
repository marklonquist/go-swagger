// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewOrderGetParams creates a new OrderGetParams object
// no default values defined in spec.
func NewOrderGetParams() OrderGetParams {

	return OrderGetParams{}
}

// OrderGetParams contains all the bound params for the order get operation
// typically these are obtained from a http.Request
//
// swagger:parameters OrderGet
type OrderGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of pet that needs to be fetched
	  Required: true
	  Maximum: 10
	  Minimum: 1
	  In: path
	*/
	OrderID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewOrderGetParams() beforehand.
func (o *OrderGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrderID, rhkOrderID, _ := route.Params.GetOK("orderId")
	if err := o.bindOrderID(rOrderID, rhkOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrderID binds and validates parameter OrderID from path.
func (o *OrderGetParams) bindOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("orderId", "path", "int64", raw)
	}
	o.OrderID = value

	if err := o.validateOrderID(formats); err != nil {
		return err
	}

	return nil
}

// validateOrderID carries on validations for parameter OrderID
func (o *OrderGetParams) validateOrderID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("orderId", "path", int64(o.OrderID), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("orderId", "path", int64(o.OrderID), 10, false); err != nil {
		return err
	}

	return nil
}
