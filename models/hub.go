package models

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (hub *Hub) NewHub() {
	hub.clients = make(map[*Client]bool)
	hub.broadcast = make(chan []byte)
	hub.register = make(chan *Client)
	hub.unregister = make(chan *Client)
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.unregister:
			delete(hub.clients, client)
		case client := <-hub.register:
			hub.clients[client] = true
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}

	}
}
