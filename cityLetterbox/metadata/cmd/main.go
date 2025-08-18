package cmd

import "cityletterbox.com/metadata/internal/repository"

func main() {
	r := repository.New() 
	c := controller.New(r)
	h := handler.New(c)
}