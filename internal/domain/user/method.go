package user

func (u *User) GetUserRoleString() string {
	return string(u.Role)
}

func (u *User) SetUserRoleString(r string) {
	switch r {
	case string(PSYCHOLOGIST):
		u.Role = PSYCHOLOGIST
	default:
		u.Role = BASE
	}
}
