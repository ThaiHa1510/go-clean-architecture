package usercasses

type RegisterUserUseCaseImpl struct {
    userRepository repository.UserRepository
}

func NewRegisterUserUseCaseImpl(userRepository repository.UserRepository) *RegisterUserUseCaseImpl {
    return &RegisterUserUseCaseImpl{
        userRepository: userRepository,
    }
}

func (uc *RegisterUserUseCaseImpl) Execute(ctx context.Context, input *RegisterUserInput) (*RegisterUserOutput, error) {
    // Validate user input
    if err := input.Validate(); err != nil {
        return nil, err
    }

    // Create a new user entity
    user := entity.NewUser(input.Name, input.Email, input.Password)

    // Persist user data to the database
    err := uc.userRepository.Create(ctx, user)
    if err != nil {
        return nil, err
    }

    // Return the newly created user's ID
    return &RegisterUserOutput{UserID: user.ID}, nil
}
