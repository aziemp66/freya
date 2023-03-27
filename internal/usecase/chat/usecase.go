package chat

import (
	"context"
	"time"

	httpCommon "github.com/aziemp66/freya-be/common/http"
)

type Usecase interface {
	InsertAppointment(ctx context.Context, pyschologistID string, userID string, date time.Time) (err error)
	FindAppointmentByID(ctx context.Context, id string) (appointment httpCommon.Appointment, err error)
	FindAppointmentByUserID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error)
	FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error)
	UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error)
	InsertChatroom(ctx context.Context, appointmentID string, psychologistID string, userID string) (err error)
	FindChatroomByID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error)
	FindChatroomByAppointmentID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error)
	DeleteChatroom(ctx context.Context, id string) (err error)
	InsertMessageToChatroom(ctx context.Context, senderId string, content string, chatRoomID string) (err error)
	FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []httpCommon.Message, err error)
	FindMessageByID(ctx context.Context, id string) (message httpCommon.Message, err error)
	DeleteMessage(ctx context.Context, id string) (err error)
}
