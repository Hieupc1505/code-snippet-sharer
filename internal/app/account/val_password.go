package account

const (
	MinPasswordLen = 6
)

type Password string

func (p Password) String() string { return string(p) }

func NewPassword(password string) (Password, error) {
	if len(password) < MinPasswordLen {
		return "", ErrPasswordInvalidLen
	}
	return Password(password), nil
}
