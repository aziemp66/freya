package mail

type (
	EmailVerification struct {
		Token string
	}

	PasswordReset struct {
		Email string
		Token string
	}
)
