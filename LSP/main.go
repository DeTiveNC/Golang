package main

import (
	"bufio"
	"github.com/detivenc/lsp-go/rpc"
	"log"
	"os"
)

const direction string = "/home/detivenc/OneDrive/Escritorio/UNI/Golang/LSP"

func main() {
	logger := GetLogger(direction + "/lsp-go.txt")
	logger.Println("Main")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		HandleMessage(logger, msg)
	}
}

func HandleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func GetLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(file, "[lsp-go]", log.Ldate|log.Lshortfile|log.Ltime)
}
