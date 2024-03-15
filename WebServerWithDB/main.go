package main

import (
	"database-example/db"
	"database-example/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer(handler *handler.EquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	handler.RegisterRoutes(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := db.InitDB()
	if database == nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	}


	equipmentHandler := handler.NewEquipmentHandler(database)

	startServer(equipmentHandler)
}
