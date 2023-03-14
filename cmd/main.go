package main

import (
	"fmt"
	"time"

	dbCommon "github.com/aziemp66/freya-be/common/db"
	httpCommon "github.com/aziemp66/freya-be/common/http"
	jwtCommon "github.com/aziemp66/freya-be/common/jwt"
	mailCommon "github.com/aziemp66/freya-be/common/mail"
	passwordCommon "github.com/aziemp66/freya-be/common/password"

	userDlv "github.com/aziemp66/freya-be/internal/delivery/user"
	userRepo "github.com/aziemp66/freya-be/internal/repository/user"
	userUc "github.com/aziemp66/freya-be/internal/usecase/user"

	"github.com/aziemp66/freya-be/common/env"

	"github.com/gin-contrib/cors"
)

func main() {
	cfg := env.LoadConfig()
	httpServer := httpCommon.NewHTTPServer(cfg.GinMode)

	db := dbCommon.NewDB(cfg.DBUrl, cfg.DBName)
	passwordManager := passwordCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.JwtSecretKey)
	mailDialer := mailCommon.New(cfg.MailEmail, cfg.MailPassword, cfg.MailHost, cfg.MailPort)

	root := httpServer.Router.Group("/api", httpCommon.MiddlewareErrorHandler())

	httpServer.Router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	userRepository := userRepo.NewUserRepositoryImplementation(db)
	userUseCase := userUc.NewUserUsecaseImplementation(userRepository, passwordManager, jwtManager, mailDialer)
	userDlv.NewUserDelivery(root, userUseCase)

	err := httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		panic(err)
	}
}
