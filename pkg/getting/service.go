package getting

import "errors"

var InvalidUserOrPwd = errors.New("invalid Username or Password")

type Service interface {
	Login(User) (UserInfo, error)
}

type Repository interface {
	FindByUsernameAndPassword(User) (UserInfo,error)
}

type service struct{
	r Repository
}

func NewService(r Repository) Service{
	return &service{r}
}

func (s *service)Login(u User) (UserInfo, error){
	return s.r.FindByUsernameAndPassword(u)
}