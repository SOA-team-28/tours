package main

import (
	"database-example/db"
	"database-example/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer() {
	database := db.InitDB()
	if database == nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	}

	router := mux.NewRouter().StrictSlash(true)

	equipmentHandler := handler.NewEquipmentHandler(database)
	equipmentHandler.RegisterRoutes(router)

	checkpointHandler := handler.NewCheckpointHandler(database)
	checkpointHandler.RegisterRoutes(router)

	tourHandler := handler.NewTourHandler(database)
	tourHandler.RegisterRoutes(router)

	mapObjectHandler := handler.NewMapObjectHandler(database)
	mapObjectHandler.RegisterRoutes(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {

	startServer()
}
