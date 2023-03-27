package client

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	jwtCommon "github.com/aziemp66/freya-be/common/jwt"
	wsCommon "github.com/aziemp66/freya-be/common/websocket"

	chatUsecase "github.com/aziemp66/freya-be/internal/usecase/chat"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs handles websocket requests from the peer.
func ServeWebSocket(ctx *gin.Context, chatUC chatUsecase.Usecase, jwtManager jwtCommon.JWTManager, roomId string) {
	fmt.Print(roomId)
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &wsCommon.Connection{Send: make(chan wsCommon.MessagePayload), Ws: ws}
	s := wsCommon.Subscription{Ctx: ctx, ChatUsecase: chatUC, JwtManager: jwtManager, Conn: c, Room: roomId}
	wsCommon.H.Register <- s
	go s.WritePump()
	go s.ReadPump()
}
