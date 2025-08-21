package cmd

import (
	"log"
	"net/http"

	"cityletterbox.com/metadata/internal/controller/metadata"
	httphandler "cityletterbox.com/metadata/internal/handler/http"
	"cityletterbox.com/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting metadata service")
	r := memory.New()
	c := metadata.New(r)
	h := httphandler.New(c)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
