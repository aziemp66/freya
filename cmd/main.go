package main

import (
	"fmt"
	"time"

	dbCommon "github.com/aziemp66/freya-be/common/db"
	httpCommon "github.com/aziemp66/freya-be/common/http"
	jwtCommon "github.com/aziemp66/freya-be/common/jwt"
	mailCommon "github.com/aziemp66/freya-be/common/mail"
	passwordCommon "github.com/aziemp66/freya-be/common/password"
	WsCommon "github.com/aziemp66/freya-be/common/websocket"

	chatDlv "github.com/aziemp66/freya-be/internal/delivery/chat"
	postDlv "github.com/aziemp66/freya-be/internal/delivery/post"
	userDlv "github.com/aziemp66/freya-be/internal/delivery/user"
	chatRepo "github.com/aziemp66/freya-be/internal/repository/chat"
	postRepo "github.com/aziemp66/freya-be/internal/repository/post"
	userRepo "github.com/aziemp66/freya-be/internal/repository/user"
	chatUc "github.com/aziemp66/freya-be/internal/usecase/chat"
	postUc "github.com/aziemp66/freya-be/internal/usecase/post"
	userUc "github.com/aziemp66/freya-be/internal/usecase/user"

	"github.com/aziemp66/freya-be/common/env"

	"github.com/gin-contrib/cors"
)

func main() {
	go WsCommon.H.Run()

	cfg := env.LoadConfig()
	httpServer := httpCommon.NewHTTPServer(cfg.GinMode)

	db := dbCommon.NewDB(cfg.DBUrl, cfg.DBName)
	passwordManager := passwordCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.JwtSecretKey)
	mailDialer := mailCommon.New(cfg.MailEmail, cfg.MailPassword, cfg.MailHost, cfg.MailPort)

	root := httpServer.Router.Group("/api", httpCommon.MiddlewareErrorHandler())

	httpServer.Router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	userRepository := userRepo.NewUserRepositoryImplementation(db)
	userUseCase := userUc.NewUserUsecaseImplementation(userRepository, passwordManager, jwtManager, mailDialer)
	userDlv.NewUserDelivery(root.Group("/user"), userUseCase, jwtManager)

	postRepository := postRepo.NewPostRepositoryImplementation(db)
	postUseCase := postUc.NewPostUsecaseImplementation(postRepository)
	postDlv.NewPostDelivery(root.Group("/post"), postUseCase, jwtManager)

	chatRepository := chatRepo.NewChatRepositoryImplementaion(db)
	chatUseCase := chatUc.NewChatUsecaseImplementation(chatRepository)
	chatDlv.NewChatDeliveryImplementation(root.Group("/chat"), chatUseCase, jwtManager)

	err := httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
