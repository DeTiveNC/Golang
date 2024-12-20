package main

import (
	"goAuth/internal/auth"
	"goAuth/internal/server"
)

func main() {

	auth.NewAuth()
	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		panic("cannot start newServer")
	}
}
