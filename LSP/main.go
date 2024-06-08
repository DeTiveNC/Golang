package main

import (
	"bufio"
	"encoding/json"
	"github.com/detivenc/lsp-go/analysis"
	"github.com/detivenc/lsp-go/lsp"
	"github.com/detivenc/lsp-go/rpc"
	"io"
	"log"
	"os"
)

const direction string = "<direction>"

func main() {
	logger := GetLogger(direction + "/lsp-go.txt")
	logger.Println("Main")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println(err)
			continue
		}
		HandleMessage(logger, writer, state, method, contents)
	}
}

func HandleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		// Reply
		msg := lsp.NewInitializeResponse(request.ID)

		WriteResponse(writer, msg)

		logger.Printf("Sent initialize response")
	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("Error unmarshalling didOpen request: %s", err)
			return
		}
		logger.Printf("Opened document: %s", notification.Params.TextDocument.URI)
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
	case "textDocument/didChange":
		var notification lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("Error unmarshalling didChange request: %s", err)
			return
		}
		logger.Printf("Changed document: %s", notification.Params.TextDocument.URI)
		for _, change := range notification.Params.ContentChanges {
			state.UpdateDocument(notification.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling hover request: %s", err)
			return
		}

		// Create a response
		reponse := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		// Send the response
		WriteResponse(writer, reponse)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling definition request: %s", err)
			return
		}

		// Create a response
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		// Send the response
		WriteResponse(writer, response)
	}
}

func WriteResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func GetLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(file, "[lsp-go]", log.Ldate|log.Lshortfile|log.Ltime)
}
