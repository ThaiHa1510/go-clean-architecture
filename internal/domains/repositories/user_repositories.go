package repositories

import (
	"fmt"
	"context"
)

type UserRepository struct{
	
}

func NewUserRepository(){
	return &UserRepository{}
}

func (repo *UserRepository) Find(ctx context.Context,id int64) entities.User, error {
	if id <= 0{
		return nil,fmt.Error("id is valid")
	}
	return entities.User{},nil
}

func (repo *UserRepository) Create(ctx context.Context, user )
