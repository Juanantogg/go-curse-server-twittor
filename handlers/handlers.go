package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Juanantogg/server-go-twittor/middlewares"
	"github.com/Juanantogg/server-go-twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registry", middlewares.CheckDB(routers.Registry)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
