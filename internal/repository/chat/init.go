package chat

import (
	"context"

	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepositoryImplementaion struct {
	db *mongo.Database
}

func NewChatRepositoryImplementaion(db *mongo.Database) *ChatRepositoryImplementaion {
	return &ChatRepositoryImplementaion{db}
}

func (c *ChatRepositoryImplementaion) InsertAppointment(ctx context.Context, appointment chatDomain.Chatroom) (err error) {
	return nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByID(ctx context.Context, id string) (appointment chatDomain.Chatroom, err error) {
	return appointment, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByUserID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	return appointments, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	return appointments, nil
}

func (c *ChatRepositoryImplementaion) UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error) {
	return nil
}

func (c *ChatRepositoryImplementaion) InsertChatroom(ctx context.Context, chatroom chatDomain.Chatroom) (err error) {
	return nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error) {
	return chatroom, nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByAppointmentID(ctx context.Context, id string) (chatrooms []chatDomain.Chatroom, err error) {
	return chatrooms, nil
}

func (c *ChatRepositoryImplementaion) InsertMessage(ctx context.Context, message chatDomain.Message) (err error) {
	return nil
}

func (c *ChatRepositoryImplementaion) FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []chatDomain.Message, err error) {
	return messages, nil
}

func (c *ChatRepositoryImplementaion) FindMessageByID(ctx context.Context, id string) (message chatDomain.Message, err error) {
	return message, nil
}

func (c *ChatRepositoryImplementaion) DeleteMessage(ctx context.Context, id string) (err error) {
	return nil
}
