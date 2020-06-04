package main

import (
	"log"
	"net/http"

	ball "github.com/chandrafortuna/assesment/domain/ball"
	h "github.com/chandrafortuna/assesment/handler"
	"github.com/gorilla/mux"
)

func main() {
	//register service
	ballRepository := ball.NewRepository([]*ball.Container{})
	ballService := ball.NewService(ballRepository)
	ballHandler := h.NewBallHandler(ballService)

	router := mux.NewRouter()
	router.HandleFunc("/balls/addToContainer", ballHandler.AddBallToContainer).Methods("POST")
	router.HandleFunc("/balls/init", ballHandler.Init).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}