package main

import (
	"github.com/a-r-g-v/kintai/app"
	"github.com/a-r-g-v/kintai/store"
	"github.com/goadesign/goa"
	"time"
)

// UserController implements the User resource.
type UserController struct {
	*goa.Controller
	db *store.DB
}

// NewUserController creates a User controller.
func NewUserController(service *goa.Service, db *store.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		db:         db,
	}
}

func ToUserMedia(user *store.UserModel) *app.User {
	return &app.User{
		ID:        user.ID,
		Href:      app.UserHref(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	user := c.db.NewUser()
	user.Name = ctx.Payload.Name
	user.CreatedAt = time.Now()

	c.db.SaveUser(user)

	// UserController_Create: end_implement

	ctx.ResponseData.Header().Set("Location", app.UserHref(user.ID))
	return ctx.Created()

}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	users, _ := c.db.GetUsers()
	res := make(app.UserCollection, len(users))
	for i, user := range users {
		res[i] = ToUserMedia(&user)

	}

	// UserController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	user, err := c.db.GetUser(ctx.UserID)
	if err != nil {
		return ctx.NotFound()
	}

	// UserController_Show: end_implement
	return ctx.OK(ToUserMedia(&user))
}
