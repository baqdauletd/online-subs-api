package main

import(
	"online-subs-api/repo"
	"online-subs-api/models"
)

func main(){
	db := repo.Connect()
	db.AutoMigrate(&models.Sub{})
}