package user

import (
	userUserCase "github.com/aziemp66/freya-be/internal/usecase/user"
)

type UserDevlivery struct {
	UserUseCase userUserCase.Usecase
}

func NewUserDelivery(userUseCase userUserCase.Usecase) *UserDevlivery {
	return &UserDevlivery{
		UserUseCase: userUseCase,
	}
}
