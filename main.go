package main

import (
	"log"
	"net/http"
	"online-subs-api/handlers"
	"online-subs-api/models"
	"online-subs-api/repo"
	"online-subs-api/router"
	"online-subs-api/services"
	"online-subs-api/utils"
)

func main(){
	utils.InitLogger()
	db := repo.Connect()
	db.AutoMigrate(&models.Sub{})

	repo := repo.NewSubsRepo(db)
	service := services.NewSubsService(repo)
	handler := handlers.NewSubHandler(service)

	mux := http.NewServeMux()
	router.Routes(mux, handler)


	log.Println("Server running at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}