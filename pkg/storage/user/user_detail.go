package user

import "golang-structure-api/pkg/adding"

type Detail struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
}

func NewDetail(user adding.User)*Detail {
	return &Detail{Username: user.Username,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
	}
}