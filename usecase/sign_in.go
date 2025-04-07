package usecase

import (
	"quizzy-classroom/provider"
)

type SignInUseCase interface {
	Execute(email string, password string) (string, error)
}

type signInUseCaseImpl struct {
	Provider *provider.Provider
}

func NewSignInUseCase(provider *provider.Provider) SignInUseCase {
	return &signInUseCaseImpl{
		Provider: provider,
	}
}

func (s *signInUseCaseImpl) Execute(email string, password string) (string, error) {
	return "", nil
}
