package response

import (
	"entities"
)
type RegisterResponse struct {
	Response
	User *entities.User
}
func(resp *Response) response(ctx *fiber.Ctx, responseCode int, message string, status string){
	return c.JSON(responseCode, &RegisterResponse{
		User user
	})
}