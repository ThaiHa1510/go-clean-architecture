package http_errors

import(
	"fmt"
)
type NotFoundError struct {
	Code int64
	Message string
}

func (err *NotFoundError) Error() string {
	return fmt.Printf("Not found error with code %d",err.Code)
}

type BadRequestError struct{
	code int64
	Message string
}

func (err *BadRequestError) Error() string {
	return fmt.Printf("Bad request error with code %d",err.Code)
}

type ForbiddenError struct {
	code int64
	Message string
}

func(err *ForbiddenError) Error() string {
	return fmt.Printf("Forbidden error with code %d",err.Code)
}

type RequestTimeout struct {
	code int64
	Message string
}

func(err *RequestTimeout) Error() string{
	return fmt.Printf("Timout error with code %d",err.Code)
}



