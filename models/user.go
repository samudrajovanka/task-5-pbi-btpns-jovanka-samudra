package models

import (
	"photo-api/app"
)

type User struct {
	app.User
	Photos []Photo `gorm:"foreignKey:UserID;contstraint:OnDelete:CASCADE" json:"photos"`
}
