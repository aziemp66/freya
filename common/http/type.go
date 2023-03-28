package http

import "time"

type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Value   interface{} `json:"value"`
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	User struct {
		Id              string    `json:"id"`
		FirstName       string    `json:"first_name"`
		LastName        string    `json:"last_name"`
		Email           string    `json:"email"`
		BirthDay        time.Time `json:"birthday"`
		Role            string    `json:"role"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		IsEmailVerified bool      `json:"is_email_verified"`
	}

	Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8,lte=104"`
	}

	AddUser struct {
		FirstName string    `json:"first_name" binding:"required"`
		LastName  string    `json:"last_name" binding:"required"`
		Email     string    `json:"email" binding:"required,email"`
		Password  string    `json:"password" binding:"required,gte=8,lte=104"`
		BirthDay  time.Time `json:"birthday" binding:"required"`
	}

	UpdateUser struct {
		FirstName string    `json:"first_name" binding:"required"`
		LastName  string    `json:"last_name" binding:"required"`
		BirthDay  time.Time `json:"birthday" binding:"required"`
	}

	UpdatePassword struct {
		OldPassword string `json:"old_password" binding:"required,gte=8,lte=104"`
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=104"`
	}

	ForgotPassword struct {
		Email string `json:"email" binding:"required,email"`
	}

	ResetPassword struct {
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=104"`
	}

	Post struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		AuthorID string `json:"author_id"`
	}

	AddPost struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		AuthorID string `json:"author_id"`
	}

	Comment struct {
		ID       string `json:"id"`
		Content  string `json:"content"`
		AuthorID string `json:"author_id"`
		PostID   string `json:"post_id"`
	}

	AddComment struct {
		Content  string `json:"content" binding:"required"`
		AuthorID string `json:"author_id"`
		PostID   string `json:"post_id"`
	}

	Appointment struct {
		Id             string    `json:"id"`
		PsychologistId string    `json:"psychologist_id"`
		UserId         string    `json:"user_id"`
		Status         string    `json:"status"`
		Date           time.Time `json:"date"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	AppointmentStatusUpdate struct {
		Status string `json:"status" binding:"required"`
	}

	Chatroom struct {
		Id             string    `json:"id"`
		AppointmentId  string    `json:"appointment_id"`
		PsychologistId string    `json:"psychologist_id"`
		UserId         string    `json:"user_id"`
		Messages       []Message `json:"messages"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	Message struct {
		Id        string    `json:"id"`
		SenderId  string    `json:"sender_id"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	AddAppointment struct {
		PsychologistId string    `json:"psychologist_id" binding:"required"`
		Date           time.Time `json:"date" binding:"required"`
	}

	AddMessage struct {
		Content    string `json:"content" binding:"required"`
		ChatroomId string `json:"chatroom_id" binding:"required"`
	}
)
