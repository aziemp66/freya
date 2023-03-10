package user

import (
	userUserCase "github.com/aziemp66/freya-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	UserUseCase userUserCase.Usecase
}

func NewUserDelivery(router *gin.RouterGroup, userUseCase userUserCase.Usecase) *UserDelivery {
	UserDelivery := &UserDelivery{
		UserUseCase: userUseCase,
	}

	return UserDelivery
}
