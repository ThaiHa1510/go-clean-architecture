package controller 

import (
	"context"
)
func RegisterUserHandler(ctx *fiber.Ctx) {
	// Decode the request body into the `RegisterRequest` struct
	var request RegisterRequest
	if err := json.NewDecoder(ctx.Body).Decode(&request); err != nil {
			// Handle decoding error
			return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
			// Handle validation errors
			return
	}

	// Proceed with user registration logic
	// ...
	registerUseCase =  usecases.NewRegisterUserCase()
	var user RegisterRequest
	user, err : registerUseCase.Execute(ctx, request)
	if err != nil{
		return http_errors.BadRequestError{Code:400,Message: fmt.Println("Failed when register user")}
	}
	return ctx.SendJson({
		status: "success",
		code: 200
	})
}
