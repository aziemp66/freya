package user

import (
	httpCommon "github.com/aziemp66/freya-be/common/http"
	"github.com/aziemp66/freya-be/common/http/middleware"
	"github.com/aziemp66/freya-be/common/jwt"
	userUserCase "github.com/aziemp66/freya-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	UserUseCase userUserCase.Usecase
	jwtManager  *jwt.JWTManager
}

func NewUserDelivery(router *gin.RouterGroup, userUseCase userUserCase.Usecase, jwtManager *jwt.JWTManager) *UserDelivery {
	UserDelivery := &UserDelivery{
		UserUseCase: userUseCase,
		jwtManager:  jwtManager,
	}

	router.POST("/login", UserDelivery.Login)
	router.POST("/register", UserDelivery.Register)
	router.POST("/forgot-password", UserDelivery.ForgotPassword)
	router.POST("/reset-password", UserDelivery.ResetPassword)

	authGroup := router.Group("/", middleware.JWTAuth(jwtManager))
	authGroup.PUT("/update", UserDelivery.Update)
	authGroup.PUT("/update-password", UserDelivery.UpdatePassword)

	return UserDelivery
}

func (u *UserDelivery) Login(c *gin.Context) {
	var loginRequest httpCommon.Login

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.Error(err)
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

	err := u.UserUseCase.Register(c, registerRequest.Email, registerRequest.Password, registerRequest.FirstName, registerRequest.LastName, registerRequest.BirthDay)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Register success",
	})
}

func (u *UserDelivery) Update(c *gin.Context) {
	var updateUserRequest httpCommon.UpdateUser

	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		return
	}

	userId := c.GetString("user_id")

	err := u.UserUseCase.Update(c, userId, updateUserRequest)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Update success",
	})
}

func (u *UserDelivery) UpdatePassword(c *gin.Context) {
	var updatePasswordRequest httpCommon.UpdatePassword

	if err := c.ShouldBindJSON(&updatePasswordRequest); err != nil {
		return
	}

	userId := c.GetString("user_id")

	user, err := u.UserUseCase.GetById(c, userId)

	if err != nil {
		c.Error(err)
		return
	}

	err = u.UserUseCase.UpdatePassword(c, user.Id, updatePasswordRequest.OldPassword, updatePasswordRequest.NewPassword)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Update password success",
	})
}

func (u *UserDelivery) ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest httpCommon.ForgotPassword

	if err := c.ShouldBindJSON(&forgotPasswordRequest); err != nil {
		return
	}

	err := u.UserUseCase.ForgotPassword(c, forgotPasswordRequest.Email)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Forgot password success",
	})
}

func (u *UserDelivery) ResetPassword(c *gin.Context) {
	var resetPasswordRequest httpCommon.ResetPassword

	if err := c.ShouldBindJSON(&resetPasswordRequest); err != nil {
		return
	}

	userId := c.Query("id")
	userToken := c.Query("token")

	err := u.UserUseCase.ResetPassword(c, userToken, userId, resetPasswordRequest.NewPassword)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Reset password success",
	})
}

func (u *UserDelivery) GetById(c *gin.Context) {
	userId := c.GetString("user_id")

	user, err := u.UserUseCase.GetById(c, userId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get user success",
		Value:   gin.H{"user": user},
	})
}
