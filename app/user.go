package app

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" valid:"alphanum,required"`
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password,omitempty" valid:"required,minstringlength(6)"`
}

type Login struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}
