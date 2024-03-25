package main

import (
	"bytes"
	"log"
	"sync"
	"text/template"
)

type Message struct {
	ClientID string
	Text     string
}

type WSMessage struct {
	Text    string      `json:"text"`
	Headers interface{} `json:"headers"`
}

type Hub struct {
	sync.RWMutex
	clients    map[*Client]bool
	messages   []*Message
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()
			log.Printf("Client %s register", client.id)
			for _, msg := range h.messages {
				client.send <- getMessageTemplate(msg)
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.Lock()
				close(client.send)
				delete(h.clients, client)
				h.Unlock()
				log.Printf("Client %s unregister", client.id)
			}
		case msg := <-h.broadcast:
			h.Lock()
			h.messages = append(h.messages, msg)
			h.Unlock()
			for client := range h.clients {
				select {
				case client.send <- getMessageTemplate(msg):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("template/messages.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	var renderedMessages bytes.Buffer
	err = tmpl.Execute(&renderedMessages, msg)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	return renderedMessages.Bytes()
}
