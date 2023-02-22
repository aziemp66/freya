package user

func (u *User) GetUserRoleString() string {
	return string(u.Role)
}

func (u *User) SetUserRoleString(r string) {
	switch r {
	case string(Admin):
		u.Role = Admin 
	default:
		u.Role = Base 
	}
}
