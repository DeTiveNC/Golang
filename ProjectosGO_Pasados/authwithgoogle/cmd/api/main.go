package main

import (
	"goAuth/internal/auth"
	"goAuth/internal/server"
)

func main() {

	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
