package chat

import (
	"context"

	httpCommon "github.com/aziemp66/freya-be/common/http"
	"github.com/aziemp66/freya-be/internal/repository/chat"
)

type ChatUsecaseImplementation struct {
	chatRepository chat.Repository
}

func NewChatUsecaseImplementation(chatRepository chat.Repository) *ChatUsecaseImplementation {
	return &ChatUsecaseImplementation{chatRepository}
}

func (c *ChatUsecaseImplementation) InsertAppointment(ctx context.Context, pyschologistID string, userID string) (err error) {
	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByID(ctx context.Context, id string) (appointment httpCommon.Appointment, err error) {
	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByUserID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error) {
	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error) {
	return
}

func (c *ChatUsecaseImplementation) UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error) {
	return
}

func (c *ChatUsecaseImplementation) InsertChatroom(ctx context.Context, appointmentID string, psychologistID string, userID string) (err error) {
	return
}

func (c *ChatUsecaseImplementation) FindChatroomByID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error) {
	return
}

func (c *ChatUsecaseImplementation) FindChatroomByAppointmentID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error) {
	return
}

func (c *ChatUsecaseImplementation) DeleteChatroom(ctx context.Context, id string) (err error) {
	return
}

func (c *ChatUsecaseImplementation) InsertMessageToChatroom(ctx context.Context, senderId string, content string, chatRoomID string) (err error) {
	return
}

func (c *ChatUsecaseImplementation) FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []httpCommon.Message, err error) {
	return
}

func (c *ChatUsecaseImplementation) FindMessageByID(ctx context.Context, id string) (message httpCommon.Message, err error) {
	return
}

func (c *ChatUsecaseImplementation) DeleteMessage(ctx context.Context, id string) (err error) {
	return
}
