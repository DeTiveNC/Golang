package main

import (
	"bufio"
	"encoding/json"
	"github.com/detivenc/lsp-go/lsp"
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
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println(err)
			continue
		}
		HandleMessage(logger, method, contents)
	}
}

func HandleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	}
}

func GetLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(file, "[lsp-go]", log.Ldate|log.Lshortfile|log.Ltime)
}
