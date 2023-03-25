package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	wsCommon "github.com/aziemp66/freya-be/common/websocket"
	wsHub "github.com/aziemp66/freya-be/common/websocket/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request, roomId string) {
	fmt.Print(roomId)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &wsHub.Connection{Send: make(chan wsCommon.MessagePayload), Ws: ws}
	s := wsHub.Subscription{Conn: c, Room: roomId}
	wsHub.H.Register <- s
	go s.WritePump()
	go s.ReadPump()
}
