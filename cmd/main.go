package main

import (
	"golang-structure-api/pkg/adding"
	"golang-structure-api/pkg/config/db"
	"golang-structure-api/pkg/getting"
	"golang-structure-api/pkg/http/rest"
	"golang-structure-api/pkg/storage/user"
	"log"
	"net/http"
)

func main() {
	dataSource := new(db.Database)
	if err := dataSource.InitDb();err!=nil{
		log.Fatal(err)
	}

	defer func(){
		err := dataSource.DB.Close()
		if err!=nil{
			panic(err)
		}
	}()

	s, _ := user.NewStorage(dataSource.DB)
	gettingService := getting.NewService(s)
	addingService := adding.NewService(s)

	router := rest.Handler(gettingService, addingService)
	log.Print("start on port 9090")
	if err := http.ListenAndServe(":9090", router);err!=nil{
		log.Fatalln(err)
	}
}
