package main

import (
	"log"

	"github.com/kristaponis/go-rest/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
