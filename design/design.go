package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("kintai", func() {
	Title("Kintai")
	Description("Attendance management system")
	Scheme("http")
	Host("localhost:8080")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})

})

var _ = Resource("health", func() {

	BasePath("/_ah")

	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})

var _ = Resource("User", func() {
	BasePath("/users")
	DefaultMedia(User)
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrive all users.")
		Response(OK, CollectionOf(User))
	})
	Action("show", func() {
		Description("Get user by id")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", Integer, "User ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {

		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/users/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})

var User = MediaType("application/vnd.user+json", func() {
	Description("A user")

	Attributes(func() {
		Attribute("id", Integer, "ID of user", func() {
			Example(1)
		})
		Attribute("href", String, "API href of user", func() {
			Example("/users/1")
		})
		Attribute("name", String, "Name of account", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")

		Required("id", "href", "name", "created_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("created_at")
	})

	View("tiny", func() {
		Description("tiny is the view used to list users")
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

})
