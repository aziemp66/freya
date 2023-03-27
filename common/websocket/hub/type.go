package hub

import (
	wsCommon "github.com/aziemp66/freya-be/common/websocket"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

	chatUC "github.com/aziemp66/freya-be/internal/usecase/chat"
)

type (
	//hub maintains the set of active Connections and broadcasts messages to the Connections.
	hub struct {
		// Registered Connections.
		Rooms map[string]map[*Connection]bool

		// Inbound messages from the Connections.
		Broadcast chan message

		// Register requests from the Connections.
		Register chan Subscription

		// Unregister requests from Connections.
		Unregister chan Subscription
	}

	// Connection is an middleman between the websocket Connection and the hub.
	Connection struct {
		// The websocket Connection.
		Ws *websocket.Conn

		// Buffered channel of outbound messages.
		Send chan wsCommon.MessagePayload
	}

	Subscription struct {
		Ctx         *gin.Context
		Conn        *Connection
		Room        string
		ChatUsecase chatUC.Usecase
	}

	message struct {
		Data wsCommon.MessagePayload
		Room string
	}
)
