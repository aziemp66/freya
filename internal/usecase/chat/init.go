package chat

import (
	"github.com/aziemp66/freya-be/internal/repository/chat"
)

type ChatUsecaseImplementation struct {
	chatRepository chat.Repository
}

func NewChatUsecaseImplementation(chatRepository chat.Repository) *ChatUsecaseImplementation {
	return &ChatUsecaseImplementation{chatRepository}
}
