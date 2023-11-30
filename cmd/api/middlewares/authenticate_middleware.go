package middleware

type AuthenticateMiddleware struct{
  skipAuthenticate bool
  type_middleware string
  authenticateService AuthenticateService
}

func (middleware *AuthenticateMiddleware) Handle(ctx context.Context, Next next) {
  if skipauthenticate  {
    next()
  }
  isAuthenticate , err := middleware.authenticateService.authenticate()
  if err != nil{
    Reponse.reponse(err)
  }
  // return 403
  
}
