package main

import (
	"log"
	"net/http"

	ball "github.com/chandrafortuna/assesment/domain/ball"
	hk "github.com/chandrafortuna/assesment/domain/homekey"
	kt "github.com/chandrafortuna/assesment/domain/kitara"
	h "github.com/chandrafortuna/assesment/handler"
	"github.com/gorilla/mux"
)

func main() {
	//register service
	ballRepository := ball.NewRepository()
	ballService := ball.NewService(ballRepository)
	ballHandler := h.NewBallHandler(ballService)

	kitaraRepository := kt.NewRepository()
	kitaraService := kt.NewService(kitaraRepository)
	kitaraHandler := h.NewKitaraHandler(kitaraService)

	homekeyService := hk.NewService()
	homekeyHandler := h.NewHomekeyHandler(homekeyService)

	router := mux.NewRouter()
	router.HandleFunc("/balls/addToContainer", ballHandler.AddBallToContainer).Methods("POST")
	router.HandleFunc("/balls/init", ballHandler.Init).Methods("POST")
	router.HandleFunc("/containers", ballHandler.GetAllContainers).Methods("GET")
	router.HandleFunc("/containers/verified", ballHandler.GetVerifiedContainer).Methods("GET")

	router.HandleFunc("/productVariant", kitaraHandler.Init).Methods("POST")
	router.HandleFunc("/productVariant", kitaraHandler.GetAll).Methods("GET")
	router.HandleFunc("/productVariant", kitaraHandler.Clear).Methods("DEL")
	router.HandleFunc("/productVariant/reserve", kitaraHandler.Reserve).Methods("POST")

	router.HandleFunc("/homekey", homekeyHandler.GetSolution).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
