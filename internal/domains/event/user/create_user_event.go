package user

import (
  "fmt"
)

type CreateUserEvent struct{
  payload : User
  name: string
}

func NewCreateUserEvent(user User) CreateUserEvent{
  return CreateUserEvent{payload: user}
}
