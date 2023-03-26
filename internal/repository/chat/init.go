package chat

import (
	"context"
	"time"

	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepositoryImplementaion struct {
	db *mongo.Database
}

func NewChatRepositoryImplementaion(db *mongo.Database) *ChatRepositoryImplementaion {
	return &ChatRepositoryImplementaion{db}
}

func (c *ChatRepositoryImplementaion) InsertAppointment(ctx context.Context, appointment chatDomain.Chatroom) (err error) {
	appointment.CreatedAt = time.Now()
	appointment.UpdatedAt = time.Now()

	_, err = c.db.Collection("appointments").InsertOne(ctx, appointment)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByID(ctx context.Context, id string) (appointment chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointment, err
	}

	err = c.db.Collection("appointments").FindOne(ctx, primitive.M{"_id": objId}).Decode(&appointment)

	if err != nil {
		return appointment, err
	}

	return appointment, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByUserID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointments, err
	}

	cursor, err := c.db.Collection("appointments").Find(ctx, primitive.M{"user_id": objId})

	if err != nil {
		return appointments, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &appointments)

	if err != nil {
		return appointments, err
	}

	return appointments, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointments, err
	}

	cursor, err := c.db.Collection("appointments").Find(ctx, primitive.M{"psychologist_id": objId})

	if err != nil {
		return appointments, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &appointments)

	if err != nil {
		return appointments, err
	}

	return appointments, nil
}

func (c *ChatRepositoryImplementaion) UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = c.db.Collection("appointments").UpdateOne(ctx, primitive.M{"_id": objId}, primitive.M{"$set": primitive.M{"status": status}})

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryImplementaion) InsertChatroom(ctx context.Context, chatroom chatDomain.Chatroom) (err error) {
	chatroom.CreatedAt = time.Now()
	chatroom.UpdatedAt = time.Now()

	_, err = c.db.Collection("chatrooms").InsertOne(ctx, chatroom)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return chatroom, err
	}

	err = c.db.Collection("chatrooms").FindOne(ctx, primitive.M{"_id": objId}).Decode(&chatroom)

	if err != nil {
		return chatroom, err
	}

	return chatroom, nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByAppointmentID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return chatroom, err
	}

	err = c.db.Collection("chatrooms").FindOne(ctx, primitive.M{"appointment_id": objId}).Decode(&chatroom)

	if err != nil {
		return chatroom, err
	}

	return chatroom, nil
}

func (c *ChatRepositoryImplementaion) DeleteChatroom(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = c.db.Collection("chatrooms").DeleteOne(ctx, primitive.M{"_id": objId})

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryImplementaion) InsertMessage(ctx context.Context, message chatDomain.Message) (err error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()

	_, err = c.db.Collection("messages").InsertOne(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []chatDomain.Message, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return messages, err
	}

	cursor, err := c.db.Collection("messages").Find(ctx, primitive.M{"chatroom_id": objId})

	if err != nil {
		return messages, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &messages)

	if err != nil {
		return messages, err
	}

	return messages, nil
}

func (c *ChatRepositoryImplementaion) FindMessageByID(ctx context.Context, id string) (message chatDomain.Message, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return message, err
	}

	err = c.db.Collection("messages").FindOne(ctx, primitive.M{"_id": objId}).Decode(&message)

	if err != nil {
		return message, err
	}

	return message, nil
}

func (c *ChatRepositoryImplementaion) DeleteMessage(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = c.db.Collection("messages").DeleteOne(ctx, primitive.M{"_id": objId})

	if err != nil {
		return err
	}

	return nil
}
