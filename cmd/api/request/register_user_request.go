package models

import (
	"fmt"
)
type RegisterUserRequest struct{
	ID        uint   `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name" validate:"required,min=3,max=255"`
	Email     string `json:"email" gorm:"column:email" validate:"required,email"`
	Salt      string `json:"salt" gorm:"column:salt"`
	Salted    string `json:"salted" gorm:"column:salted"`
	IconImage string `json:"icon_image" gorm:"column:icon_image"`
	Age uint16 `json:age gorm:"column:age" validate:"required,min=10, max=150"`
	Birthday string `json:birthday`
}

type (req *RegisterUserRequest) Messages() map[string]string{
	return map[string]string{
		"age.min": "Age must greater than 18 %d",
		"age.max": "Age must less than %d"
	}
}

type (req *RegisterUserRequest) Rules() map[string]string{
	return map[string]string{
		"age": "min.18",
		"age": "max.150"
	}
}

func (request *RegisterRequest) Validate() error {
	if err := validate.Struct(request); err != nil {
			return err
	}

	return nil
}