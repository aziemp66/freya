package user

func (u *User) GetUserRoleString() string {
	return string(u.Role)
}

func (u *User) SetUserRoleString(r string) {
	switch r {
	case string(Psychologist):
		u.Role = Psychologist
	default:
		u.Role = Base
	}
}
