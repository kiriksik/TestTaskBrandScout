package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kiriksik/TestTaskBrandScout/internal/config"
	"github.com/kiriksik/TestTaskBrandScout/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed loading enviroment: %s", err)
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("failed to load port")
	}

	cfg := config.InitializeApiConfig()

	serveMux := handler.InitializeMux(cfg)
	httpServer := http.Server{
		Addr:    port,
		Handler: serveMux,
	}
	fmt.Println("server started")
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}

}
