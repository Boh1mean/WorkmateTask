package main

import (
	"log"
	"net/http"

	"github.com/Boh1mean/workmateTask/internal/service"
	"github.com/Boh1mean/workmateTask/internal/transport"
)

func main() {
	store := service.NewMemoryStorage()

	usecase := service.NewTaskService(store)

	handler := transport.NewHandler(usecase)

	router := transport.NewRouter(handler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
