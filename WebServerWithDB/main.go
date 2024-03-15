package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=super dbname=tours-microservice port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.Checkpoint{}, &model.Equipment{}, &model.Tour{})
	if err != nil {
		print(err)
		return nil
	}

	return database
}

func startServer(handler *handler.EquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/equipments/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/equipments", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	startServer(equipmentHandler)
}
