package account

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email Email, password Password) *User {
	return &User{Email: email.String(), Password: password.String()}
}
