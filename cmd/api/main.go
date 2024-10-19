package main

import (
	"go/stripe/configs"
	"go/stripe/internal/server"
)

func main() {

	config.LoadConfig()
	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
