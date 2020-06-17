package adding

import "errors"

var UsernameIsRequired = errors.New("username is required")
var PasswordIsRequired = errors.New("password is required")

type Service interface {
	Register(user User) error
}

type Repository interface {
	Save(User) error
}

type service struct{
	r Repository
}

func NewService(r Repository) Service{
	return &service{r}
}

func (s *service)Register(user User) error{
	if user.Username == ""{
		return UsernameIsRequired
	}
	if user.Password == ""{
		return PasswordIsRequired
	}
	return s.r.Save(user)
}