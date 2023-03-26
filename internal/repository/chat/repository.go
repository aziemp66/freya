package chat

import (
	"context"

	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"
)

type Repository interface {
	InsertAppointment(ctx context.Context, appointment chatDomain.Appointment) (err error)
	FindAppointmentByID(ctx context.Context, id string) (appointment chatDomain.Appointment, err error)
	FindAppointmentByUserID(ctx context.Context, id string) (appointments []chatDomain.Appointment, err error)
	FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []chatDomain.Appointment, err error)
	UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error)
	InsertChatroom(ctx context.Context, chatroom chatDomain.Chatroom) (err error)
	FindChatroomByID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error)
	FindChatroomByAppointmentID(ctx context.Context, id string) (chatrooms chatDomain.Chatroom, err error)
	DeleteChatroom(ctx context.Context, id string) (err error)
	InsertMessageToChatroom(ctx context.Context, message chatDomain.Message, chatroomId string) (err error)
	FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []chatDomain.Message, err error)
	FindMessageByID(ctx context.Context, id string) (message chatDomain.Message, err error)
	DeleteMessage(ctx context.Context, id string) (err error)
}
