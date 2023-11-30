package services

import (
  "fmt"
)

type AuthenticateServiceInterface interface{
   Authenticate() *User,error
   GetUserInfo(int64 id) *User,error
   IsValid(string token) bool
}

type Calm struct {
  data map(string,interface{})
}

type AuthenticateService struct{
  keySercet string
  calm Calm
}
