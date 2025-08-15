package main

import (
	"context"
	"iLikeToKnow.com/internal/initializers"
	"log"
	"net/http"

	httpapi "iLikeToKnow.com/internal/http"
)

var domainServiceInitializerFunc = initializers.NewDefaultDomainService

func main() {
	ctx := context.Background()

	deps, err := domainServiceInitializerFunc(ctx)
	if err != nil {
		log.Fatal("domain initializer failed")
	}

	handler := httpapi.NewEventsHandler(deps.DomainService)

	http.HandleFunc("/events", handler.CreateEvent)   // POST /events
	http.HandleFunc("/events/", handler.GetEventById) // GET /events/{id}

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
