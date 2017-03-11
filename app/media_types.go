// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/a-r-g-v/kintai/design
// --out=$(GOPATH)src/github.com/a-r-g-v/kintai
// --version=v1.1.0-dirty
//
// API "kintai": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// A user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID int `form:"id" json:"id" xml:"id"`
	// Name of account
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// A user (link view)
//
// Identifier: application/vnd.user+json; view=link
type UserLink struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// A user (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID int `form:"id" json:"id" xml:"id"`
	// Name of account
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (link view)
//
// Identifier: application/vnd.user+json; type=collection; view=link
type UserLinkCollection []*UserLink

// Validate validates the UserLinkCollection media type instance.
func (mt UserLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (tiny view)
//
// Identifier: application/vnd.user+json; type=collection; view=tiny
type UserTinyCollection []*UserTiny

// Validate validates the UserTinyCollection media type instance.
func (mt UserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}