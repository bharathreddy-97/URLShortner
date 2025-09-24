package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httphandlers "github.com/bharathreddy-97/URLShortner/Handlers"
	persistence "github.com/bharathreddy-97/URLShortner/Persistence"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	dataMap := persistence.GetSMap()
	initailizeServer(dataMap)
}

func initailizeServer(dataMap *persistence.SMap) {
	mux := http.NewServeMux()

	handlers := httphandlers.Handlers{MapData: dataMap}

	mux.HandleFunc("/add", handlers.ShortenURL)
	mux.HandleFunc("/metrics", handlers.GetMetrics)
	mux.HandleFunc("/{id}", handlers.RedirectURL)

	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("BASE_URL"), os.Getenv("PORT")), httphandlers.Middleware(mux))
}
