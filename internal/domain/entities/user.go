package entities

import (
	"fmt"
)

type User struct{
	Id int64
	FristName string
	LastName string
	Age uint16
	Address string
}

func (user *User) GetFullName() string{
	return  fmt.Printf("%s %s",user.FristName,user.LastName)
}
