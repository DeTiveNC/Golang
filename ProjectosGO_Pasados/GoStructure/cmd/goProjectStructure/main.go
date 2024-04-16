package main

import (
	"fmt"
	"net/http"

	"github.com/DeTiveNC/goProjectStructure/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8081
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
