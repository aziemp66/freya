package user

import (
	"github.com/aziemp66/freya-be/internal/domain"
)

type (
	User struct {
		ID              string
		FirstName       string
		LastName        string
		Email           string
		Password        string
		IsEmailVerified bool
		Role            role

		Timestamp
	}

	role      string
	Timestamp = domain.Timestamp
)
