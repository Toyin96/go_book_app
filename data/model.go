package data

import (
	"image"
	"time"
)

type Book struct {
	Title     string
	Text      string
	Image     image.Image
	Author    User
	DateAdded time.Time
}

type User struct {
	FirstName string
	LastName string
	Picture image.Image
	EmailAddress string
}
