package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/YaelJBS/Practicing_Golang/middlew"
	"github.com/YaelJBS/Practicing_Golang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", middlew.CheckDB(routers.ToPost)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
