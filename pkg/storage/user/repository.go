package user

import (
	"github.com/jinzhu/gorm"
	"golang-structure-api/pkg/adding"
	"golang-structure-api/pkg/getting"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) (*Storage, error){
	s := new(Storage)
	s.db = db
	return s, nil
}

func (s *Storage)FindByUsernameAndPassword(user getting.User) (getting.UserInfo,error){
	d := new(Detail)
	var userInfo getting.UserInfo

	result := s.db.Take(&Detail{Username: user.Username, Password: user.Password}).Scan(&d)
	if result.Error!=nil{
		return userInfo, result.Error
	}
	if result.RecordNotFound(){
		return userInfo, getting.InvalidUserOrPwd
	}

	userInfo.FirstName = d.FirstName
	userInfo.LastName = d.LastName
	userInfo.Age = d.Age
	return userInfo, nil
}

func (s *Storage)Save(user adding.User) error {
	result := s.db.Debug().Create(NewDetail(user))
	if result.Error!=nil{
		return result.Error
	}
	return nil
}