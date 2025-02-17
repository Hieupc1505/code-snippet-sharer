package account

import (
	"context"
	"s-coder-snippet-sharder/pkg/errsx"
)

func signupUser(ctx context.Context, email, password string) (*User, error) {
	var errs errsx.Map
	pwd, err := NewPassword(password)
	if err != nil {
		errs.Set("password", err)
	}

	emailAddr, err := NewEmail(email)
	if err != nil {
		errs.Set("email", err)
	}

	if errs != nil {
		return nil, err
	}

	user := NewUser(emailAddr, pwd)
	return user, nil
}

func (s *Service) SignUp(ctx context.Context, email, password string) (*User, error) {
	user, err := signupUser(ctx, email, password)
	if err != nil {
		return nil, err
	}
	//user.SignedUp()
	//Todo: Store hashed password
	//if err := s.repo.AddUser(ctx, user); err != nil {
	//	return nil, err
	//}
	return user, nil
}
