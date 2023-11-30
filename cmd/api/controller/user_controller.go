package controller

type UserController struct {
  basePath string
  data map(string, interface{})
  userService UserService
}

func NewUserController(UserService userService ,string basePath) UserController {
  return UserController{
    userService: userService
    basePath: basePath
  }
} 

func(controller *UserController) Get(ctx context.Context, interface{} request) Reponse, error{
    user , err := controller.userService.Find(request['id'])
    if err != nil{
      return nil, err
    }
    if user == nil{
      return nil,common.ErrorRequestNotFound
    }
    return UserReponse{ user: user}
}
