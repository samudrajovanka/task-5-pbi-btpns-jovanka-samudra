package models

import (
	"photo-api/app"
)

type Photo struct {
	app.Photo
	UserID int
}
