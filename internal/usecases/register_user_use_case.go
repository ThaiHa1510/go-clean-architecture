package usecases

import (
	"context"
	"api"
	"events"
	"listener"
)
type RegisterUserUseCaseImpl struct{
	userRepo repository.UserRepository
	ctx context.Context
}

func 

func(regUserCase *RegisterUserUseCaseImpl) Execute(ctx  context.Context, input api.RegisterUserInput) api.RegeisterUserOuput,error {
		user , err :=regUserCase.userRepo.Create(input)
		if err != nil{
			return nil, err
		}
		listener := NewRegisterUserListener()
		event :=  events.NewRegisterUserEvent(events.WithNameOption("name"),events.WithRetry(3))
		event.Register(listener)
		event.NotifyAll()
		return &RegeisterUserOuput{Id:user.Id},nil
}