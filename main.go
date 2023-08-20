package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Port: ", port)

	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders: []string{"Link"},
			MaxAge:         300,
		}),
	)

	v1router := chi.NewRouter()
	v1router.Get("/ready", handleReadiness)
	v1router.Get("/err", handleErr)
	// v1router.Get("/ready", func(writer http.ResponseWriter, request *http.Request) {

	router.Mount("/v1", v1router)

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Printf("server started at %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
