package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-structure-api/pkg/storage/user"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database)InitDb() error{
	db, err := gorm.Open("mysql", "root:P@ssw0rd@tcp(192.168.1.33:33060)/user_info?charset=utf8&parseTime=True")
	if err!=nil{
		return err
	}
	db.SingularTable(true)
	db.AutoMigrate(&user.Detail{})
	d.DB = db
	return nil
}