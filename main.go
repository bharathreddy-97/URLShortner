package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httphandlers "github.com/bharathreddy-97/URLShortner/Handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	initailizeServer()
}

func initailizeServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", httphandlers.ShortenURL)
	mux.HandleFunc("/metrics", httphandlers.GetMetrics)
	mux.HandleFunc("/{id}", httphandlers.RedirectURL)

	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("BASE_URL"), os.Getenv("PORT")), httphandlers.Middleware(mux))
}
