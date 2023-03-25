package chat

import (
	"time"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
)

type (
	Chatroom struct {
		ID           string          `bson:"_id,omitempty"`
		Psychologist userDomain.User `bson:"psychologist,omitempty"`
		User         userDomain.User `bson:"user,omitempty"`
		ChatHistory  []Message       `bson:"chat_history,omitempty"`
		CreatedAt    time.Time       `bson:"created_at,omitempty"`
		UpdatedAt    time.Time       `bson:"updated_at,omitempty"`
	}

	Message struct {
		SenderId  string    `bson:"sender_id,omitempty"`
		Content   string    `bson:"content,omitempty"`
		CreatedAt time.Time `bson:"created_at,omitempty"`
		UpdatedAt time.Time `bson:"updated_at,omitempty"`
	}
)
