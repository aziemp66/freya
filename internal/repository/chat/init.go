package chat

import (
	"context"
	"time"

	errorCommon "github.com/aziemp66/freya-be/common/error"
	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"

	"go.mongodb.org/mongo-driver/bson"
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
		return errorCommon.NewInvariantError("Failed to insert appointment")
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByID(ctx context.Context, id string) (appointment chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointment, errorCommon.NewInvariantError("Invalid appointment id format")
	}

	err = c.db.Collection("appointments").FindOne(ctx, bson.M{"_id": objId}).Decode(&appointment)

	if err != nil {
		return appointment, errorCommon.NewInvariantError("Appointment not found")
	}

	return appointment, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByUserID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Invalid user id format")
	}

	cursor, err := c.db.Collection("appointments").Find(ctx, bson.M{"user_id": objId})

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Appointment not found")
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &appointments)

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Appointment not found")
	}

	return appointments, nil
}

func (c *ChatRepositoryImplementaion) FindAppointmentByPsychologistID(ctx context.Context, id string) (appointments []chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Invalid psychologist id format")
	}

	cursor, err := c.db.Collection("appointments").Find(ctx, bson.M{"psychologist_id": objId})

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Appointment not found")
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &appointments)

	if err != nil {
		return appointments, errorCommon.NewInvariantError("Appointment not found")
	}

	return appointments, nil
}

func (c *ChatRepositoryImplementaion) UpdateAppointmentStatus(ctx context.Context, id string, status string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid appointment id format")
	}

	_, err = c.db.Collection("appointments").UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": bson.M{"status": status}})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to update appointment status")
	}

	return nil
}

func (c *ChatRepositoryImplementaion) InsertChatroom(ctx context.Context, chatroom chatDomain.Chatroom) (err error) {
	chatroom.CreatedAt = time.Now()
	chatroom.UpdatedAt = time.Now()

	_, err = c.db.Collection("chatrooms").InsertOne(ctx, chatroom)

	if err != nil {
		return errorCommon.NewInvariantError("Failed to insert chatroom")
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return chatroom, errorCommon.NewInvariantError("Invalid chatroom id format")
	}

	err = c.db.Collection("chatrooms").FindOne(ctx, bson.M{"_id": objId}).Decode(&chatroom)

	if err != nil {
		return chatroom, errorCommon.NewInvariantError("Chatroom not found")
	}

	return chatroom, nil
}

func (c *ChatRepositoryImplementaion) FindChatroomByAppointmentID(ctx context.Context, id string) (chatroom chatDomain.Chatroom, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return chatroom, errorCommon.NewInvariantError("Invalid appointment id format")
	}

	err = c.db.Collection("chatrooms").FindOne(ctx, bson.M{"appointment_id": objId}).Decode(&chatroom)

	if err != nil {
		return chatroom, errorCommon.NewInvariantError("Chatroom not found")
	}

	return chatroom, nil
}

func (c *ChatRepositoryImplementaion) DeleteChatroom(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid chatroom id format")
	}

	_, err = c.db.Collection("chatrooms").DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to delete chatroom")
	}

	return nil
}

func (c *ChatRepositoryImplementaion) InsertMessageToChatrooms(ctx context.Context, message chatDomain.Message, chatroomId string) (err error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()

	objId, err := primitive.ObjectIDFromHex(chatroomId)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid chatroom id format")
	}

	_, err = c.db.Collection("chatrooms").UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$push": bson.M{"messages": message}})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to insert message to chatroom")
	}

	return nil
}

func (c *ChatRepositoryImplementaion) FindAllMessagesByChatroomID(ctx context.Context, id string) (messages []chatDomain.Message, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return messages, errorCommon.NewInvariantError("Invalid chatroom id format")
	}

	var chatroom chatDomain.Chatroom

	err = c.db.Collection("chatrooms").FindOne(ctx, bson.M{"_id": objId}).Decode(&chatroom)

	if err != nil {
		return messages, errorCommon.NewInvariantError("Chatroom not found")
	}

	messages = chatroom.Messages

	return messages, nil
}

func (c *ChatRepositoryImplementaion) FindMessageByID(ctx context.Context, id string) (message chatDomain.Message, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return message, errorCommon.NewInvariantError("Invalid message id format")
	}

	var chatroom chatDomain.Chatroom

	err = c.db.Collection("chatrooms").FindOne(ctx, bson.D{{Key: "messages", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "_id", Value: objId}}}}}}).Decode(&chatroom)

	if err != nil {
		return message, errorCommon.NewInvariantError("Message not found")
	}

	for _, msg := range chatroom.Messages {
		if msg.ID == objId {
			message = msg
		}
	}

	return message, nil
}

func (c *ChatRepositoryImplementaion) DeleteMessage(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid message id format")
	}

	_, err = c.db.Collection("chatrooms").UpdateOne(ctx, bson.D{{Key: "messages", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "_id", Value: objId}}}}}}, bson.M{"$pull": bson.M{"messages": bson.M{"_id": objId}}})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to delete message")
	}

	return nil
}
