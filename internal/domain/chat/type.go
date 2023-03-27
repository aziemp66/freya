package chat

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Appointment struct {
		ID             primitive.ObjectID `bson:"_id,omitempty"`
		PsychologistID primitive.ObjectID `bson:"psychologist_id,omitempty"`
		UserID         primitive.ObjectID `bson:"user_id,omitempty"`
		Status         string             `bson:"status,omitempty"`
		Date           time.Time          `bson:"date,omitempty"`
		CreatedAt      time.Time          `bson:"created_at,omitempty"`
		UpdatedAt      time.Time          `bson:"updated_at,omitempty"`
	}

	Chatroom struct {
		ID             primitive.ObjectID `bson:"_id,omitempty"`
		AppointmentID  primitive.ObjectID `bson:"appointment_id,omitempty"`
		PsychologistID primitive.ObjectID `bson:"psychologist_id,omitempty"`
		UserID         primitive.ObjectID `bson:"user_id,omitempty"`
		Messages       []Message          `bson:"messages,omitempty"`
		CreatedAt      time.Time          `bson:"created_at,omitempty"`
		UpdatedAt      time.Time          `bson:"updated_at,omitempty"`
	}

	Message struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		SenderID  primitive.ObjectID `bson:"sender_id,omitempty"`
		Content   string             `bson:"content,omitempty"`
		CreatedAt time.Time          `bson:"created_at,omitempty"`
		UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	}
)
