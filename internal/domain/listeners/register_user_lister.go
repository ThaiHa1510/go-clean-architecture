package listener

import (
	"fmt"
)
type RegisterUserListener struct {
	userRepo UserRepository
}

func(listener * RegisterUserListener) Excute(event Event){
	fmt.Println("Excute run")
}