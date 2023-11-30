package listeners

import (
  "fmt"
)

type UserListener struct{
  emailService: EmailService
}

func(listener *UserListener) OnCreateUserListener(event CreateUserEvent) error{
   payload,err := event.GetPayload()
  if err != nil{
    return err
  }
  return nil
}
