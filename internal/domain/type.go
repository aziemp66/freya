package domain

import "time"

type (
	Timestamp struct {
		CreatedAt time.Time `bson:"created_at"`
		UpdatedAt time.Time `bson:"updated_at"`
	}
)
