package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"

	errorCommon "github.com/aziemp66/freya-be/common/error"
)

var H = hub{
	Broadcast:  make(chan message),
	Register:   make(chan Subscription),
	Unregister: make(chan Subscription),
	Rooms:      make(map[string]map[*Connection]bool),
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.Register:
			Connections := h.Rooms[s.Room]
			if Connections == nil {
				Connections = make(map[*Connection]bool)
				h.Rooms[s.Room] = Connections
			}
			h.Rooms[s.Room][s.Conn] = true
		case s := <-h.Unregister:
			Connections := h.Rooms[s.Room]
			if Connections != nil {
				if _, ok := Connections[s.Conn]; ok {
					delete(Connections, s.Conn)
					close(s.Conn.Send)
					if len(Connections) == 0 {
						delete(h.Rooms, s.Room)
					}
				}
			}
		case m := <-h.Broadcast:
			Connections := h.Rooms[m.Room]
			for c := range Connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(Connections, c)
					if len(Connections) == 0 {
						delete(h.Rooms, m.Room)
					}
				}
			}
		}
	}
}

// readPump pumps messages from the websocket Connection to the hub.
func (s Subscription) ReadPump() {
	c := s.Conn
	defer func() {
		H.Unregister <- s
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(MaxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(PongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		var readPayload ReadPayload

		err := c.Ws.ReadJSON(&readPayload)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err.Error())
			}
			break
		}

		user := s.Ctx.GetString("user_id")

		messagePayload := MessagePayload{
			User:    user,
			Message: readPayload.Message,
		}

		m := message{messagePayload, s.Room}

		H.Broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(payload MessagePayload) error {
	c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
	return c.Ws.WriteJSON(payload)
}

// writePump pumps messages from the hub to the websocket Connection.
func (s *Subscription) WritePump() {
	c := s.Conn
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(MessagePayload{})
				return
			}

			err := s.ChatUsecase.InsertMessageToChatroom(s.Ctx, message.User, message.Message, s.Room)

			if err != nil {
				fmt.Printf("error: %v", err.Error())
				s.Ctx.AbortWithError(500, errorCommon.NewInvariantError(err.Error()))
				return
			}

			if err := c.write(message); err != nil {
				s.Ctx.AbortWithError(500, errorCommon.NewInvariantError(err.Error()))
				return
			}
		case <-ticker.C:
			if err := c.write(MessagePayload{}); err != nil {
				log.Fatalf("error: %v", err)
				return
			}
		}
	}
}
