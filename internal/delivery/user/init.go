package user

import (
	httpCommon "github.com/aziemp66/freya-be/common/http"
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

	router.POST("/login", UserDelivery.Login)
	router.POST("/register", UserDelivery.Register)

	return UserDelivery
}

func (u *UserDelivery) Login(c *gin.Context) {
	var loginRequest httpCommon.Login

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		return
	}

	token, err := u.UserUseCase.Login(c, loginRequest.Email, loginRequest.Password)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Login success",
		Value:   gin.H{"token": token},
	})
}

func (u *UserDelivery) Register(c *gin.Context) {
	var registerRequest httpCommon.AddUser

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		return
	}

	err := u.UserUseCase.Register(c, registerRequest)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Register success",
	})
}
