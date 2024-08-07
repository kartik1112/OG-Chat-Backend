package models

import "fmt"

type Hub struct {
	clients       map[*Client]bool
	broadcast     chan []byte
	register      chan *Client
	unregister    chan *Client
	clientDetails map[string]*Client
}

func (hub *Hub) NewHub() {
	hub.clients = make(map[*Client]bool)
	hub.broadcast = make(chan []byte)
	hub.register = make(chan *Client)
	hub.unregister = make(chan *Client)
	hub.clientDetails = make(map[string]*Client)
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.unregister:
			delete(hub.clients, client)
		case client := <-hub.register:
			hub.clients[client] = true
			hub.clientDetails[client.selfEmail] = client
		case message := <-hub.broadcast:
			for client := range hub.clients {
				client.send <- message
			}
		}

	}
}

func (hub *Hub) SendDirectMessage(to string, message []byte) {
	fmt.Print(to)
	// client := hub.clientDetails
	if client, ok := hub.clientDetails[to]; ok {
		client.send <- message
	}

}
