//go:generate goagen bootstrap -d github.com/a-r-g-v/kintai/design

package main

import (
	"github.com/a-r-g-v/kintai/app"
	"github.com/a-r-g-v/kintai/store"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("kintai")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db := store.NewDB("./test.db")
	// Mount "User" controller
	c := NewUserController(service, db)
	app.MountUserController(service, c)
	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
