package main

import (
	"log"
	"net/http"

	"cityletterbox.com/rating/internal/controller/rating"
	httpHandler "cityletterbox.com/rating/internal/handler/http"
	"cityletterbox.com/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
