package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/detivenc/websockets0/internal/socket"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go socket.Writer(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	fmt.Println("Youtube Subscriber Monitor")

	setUpRoutes()
}

// Nota mental acordarme que siempre las funciones la inicial en mayuscula al contrario de java
