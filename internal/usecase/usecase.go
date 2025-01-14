package usecase

type UseCases struct {
	UserUseCase *UserUseCase
}

func LoadUseCases() *UseCases {
	return &UseCases{}
}
