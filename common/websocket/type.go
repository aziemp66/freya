package websocket

type (
	// WebSocket Payload
	MessagePayload struct {
		User    string `json:"user"`
		Message string `json:"message"`
	}
	ReadPayload struct {
		Message string `json:"message"`
	}
)
