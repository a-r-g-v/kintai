// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/a-r-g-v/kintai/design
// --out=$(GOPATH)src/github.com/a-r-g-v/kintai
// --version=v1.1.0-dirty
//
// API "kintai": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

// CreateUserContext provides the User create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller create action.
func NewCreateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createUserPayload is the User create action payload.
type createUserPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateUserPayload is the User create action payload.
type CreateUserPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// ListUserContext provides the User list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller list action.
func NewListUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	if r == nil {
		r = UserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKLink(r UserLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	if r == nil {
		r = UserLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKTiny(r UserTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	if r == nil {
		r = UserTinyCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowUserContext provides the User show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller show action.
func NewShowUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
		if rctx.UserID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`userId`, rctx.UserID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKLink(r *UserLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKTiny(r *UserTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// HealthHealthContext provides the health health action context.
type HealthHealthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHealthHealthContext parses the incoming request URL and body, performs validations and creates the
// context used by the health controller health action.
func NewHealthHealthContext(ctx context.Context, r *http.Request, service *goa.Service) (*HealthHealthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := HealthHealthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *HealthHealthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}