package chat

import (
	"context"
	"time"

	httpCommon "github.com/aziemp66/freya-be/common/http"
	"go.mongodb.org/mongo-driver/bson/primitive"

	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"
	chatRepository "github.com/aziemp66/freya-be/internal/repository/chat"
)

type ChatUsecaseImplementation struct {
	chatRepository chatRepository.Repository
}

func NewChatUsecaseImplementation(chatRepository chatRepository.Repository) *ChatUsecaseImplementation {
	return &ChatUsecaseImplementation{chatRepository}
}

func (c *ChatUsecaseImplementation) InsertAppointment(ctx context.Context, pyschologistID string, userID string, date time.Time) (err error) {
	psyObjID, err := primitive.ObjectIDFromHex(pyschologistID)
	if err != nil {
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return
	}

	appointment := chatDomain.Appointment{
		ID:             primitive.NewObjectID(),
		PsychologistID: psyObjID,
		UserID:         userObjID,
		Status:         httpCommon.APPOINTMENTPENDING,
		Date:           date,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = c.chatRepository.InsertAppointment(ctx, appointment)

	if err != nil {
		return err
	}

	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByID(ctx context.Context, id string) (appointment httpCommon.Appointment, err error) {
	app, err := c.chatRepository.FindAppointmentByID(ctx, id)

	if err != nil {
		return appointment, err
	}

	appointment = httpCommon.Appointment{
		Id:             app.ID.Hex(),
		PsychologistId: app.PsychologistID.Hex(),
		UserId:         app.UserID.Hex(),
		Status:         app.Status,
		Date:           app.Date,
		CreatedAt:      app.CreatedAt,
		UpdatedAt:      app.UpdatedAt,
	}

	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByUserID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error) {
	apps, err := c.chatRepository.FindAppointmentByUserID(ctx, id)

	if err != nil {
		return appointments, err
	}

	for _, app := range apps {
		appointment := httpCommon.Appointment{
			Id:             app.ID.Hex(),
			PsychologistId: app.PsychologistID.Hex(),
			UserId:         app.UserID.Hex(),
			Status:         app.Status,
			Date:           app.Date,
			CreatedAt:      app.CreatedAt,
			UpdatedAt:      app.UpdatedAt,
		}

		appointments = append(appointments, appointment)
	}

	return
}

func (c *ChatUsecaseImplementation) FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []httpCommon.Appointment, err error) {
	apps, err := c.chatRepository.FindAppointmentByPsychologistID(ctx, id)

	if err != nil {
		return appointments, err
	}

	for _, app := range apps {
		appointment := httpCommon.Appointment{
			Id:             app.ID.Hex(),
			PsychologistId: app.PsychologistID.Hex(),
			UserId:         app.UserID.Hex(),
			Status:         app.Status,
			Date:           app.Date,
			CreatedAt:      app.CreatedAt,
			UpdatedAt:      app.UpdatedAt,
		}

		appointments = append(appointments, appointment)
	}

	return
}

func (c *ChatUsecaseImplementation) UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error) {
	// if status is not equal to predefined status, return error
	if status != httpCommon.APPOINTMENTPENDING &&
		status != httpCommon.APPOINTMENTACCEPTED &&
		status != httpCommon.APPOINTMENTREJECTED &&
		status != httpCommon.APPOINTMENTCOMPLETED &&
		status != httpCommon.APPOINTMENTCANCELED {
		return
	}

	err = c.chatRepository.UpdateAppointmentStatus(ctx, id, status)

	if err != nil {
		return err
	}

	return
}

func (c *ChatUsecaseImplementation) InsertChatroom(ctx context.Context, appointmentID string, psychologistID string, userID string) (err error) {
	appointmentObjID, err := primitive.ObjectIDFromHex(appointmentID)
	if err != nil {
		return
	}

	psychologistObjID, err := primitive.ObjectIDFromHex(psychologistID)
	if err != nil {
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return
	}

	chatroom := chatDomain.Chatroom{
		ID:             primitive.NewObjectID(),
		AppointmentID:  appointmentObjID,
		PsychologistID: psychologistObjID,
		UserID:         userObjID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = c.chatRepository.InsertChatroom(ctx, chatroom)

	return
}

func (c *ChatUsecaseImplementation) FindChatroomByID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error) {
	chat, err := c.chatRepository.FindChatroomByID(ctx, id)

	if err != nil {
		return chatroom, err
	}

	chatroom = httpCommon.Chatroom{
		Id:             chat.ID.Hex(),
		AppointmentId:  chat.AppointmentID.Hex(),
		PsychologistId: chat.PsychologistID.Hex(),
		UserId:         chat.UserID.Hex(),
		CreatedAt:      chat.CreatedAt,
		UpdatedAt:      chat.UpdatedAt,
	}

	return
}

func (c *ChatUsecaseImplementation) FindChatroomByAppointmentID(ctx context.Context, id string) (chatroom httpCommon.Chatroom, err error) {
	chat, err := c.chatRepository.FindChatroomByAppointmentID(ctx, id)

	if err != nil {
		return chatroom, err
	}

	chatroom = httpCommon.Chatroom{
		Id:             chat.ID.Hex(),
		AppointmentId:  chat.AppointmentID.Hex(),
		PsychologistId: chat.PsychologistID.Hex(),
		UserId:         chat.UserID.Hex(),
		CreatedAt:      chat.CreatedAt,
		UpdatedAt:      chat.UpdatedAt,
	}

	return
}

func (c *ChatUsecaseImplementation) DeleteChatroom(ctx context.Context, id string) (err error) {
	err = c.chatRepository.DeleteChatroom(ctx, id)

	return
}

func (c *ChatUsecaseImplementation) InsertMessageToChatroom(ctx context.Context, senderId string, content string, chatRoomID string) (err error) {
	senderObjID, err := primitive.ObjectIDFromHex(senderId)
	if err != nil {
		return
	}

	message := chatDomain.Message{
		ID:        primitive.NewObjectID(),
		SenderID:  senderObjID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = c.chatRepository.InsertMessageToChatroom(ctx, message, chatRoomID)

	return
}

func (c *ChatUsecaseImplementation) FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []httpCommon.Message, err error) {
	msgs, err := c.chatRepository.FindAllMessagesByChatroomID(ctx, id)

	if err != nil {
		return messages, err
	}

	for _, msg := range msgs {
		message := httpCommon.Message{
			Id:        msg.ID.Hex(),
			SenderId:  msg.SenderID.Hex(),
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt,
			UpdatedAt: msg.UpdatedAt,
		}

		messages = append(messages, message)
	}

	return
}

func (c *ChatUsecaseImplementation) FindMessageByID(ctx context.Context, id string) (message httpCommon.Message, err error) {
	msg, err := c.chatRepository.FindMessageByID(ctx, id)

	if err != nil {
		return message, err
	}

	message = httpCommon.Message{
		Id:        msg.ID.Hex(),
		SenderId:  msg.SenderID.Hex(),
		Content:   msg.Content,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
	}

	return
}

func (c *ChatUsecaseImplementation) DeleteMessage(ctx context.Context, id string) (err error) {
	err = c.chatRepository.DeleteMessage(ctx, id)

	return
}
