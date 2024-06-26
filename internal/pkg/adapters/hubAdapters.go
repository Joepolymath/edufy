package adapters

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"go.uber.org/zap"
)

type HubInterface interface {
	Register(client *websocket.Conn, route string)
	Unregister(client *websocket.Conn, route string)
	Broadcast(client *websocket.Conn, route string, message []byte)
	Run()
}
type Hub struct {
	clients    map[*websocket.Conn]map[string]bool // Map of clients per route
	broadcast  chan Message
	register   chan RegisterMessage
	unregister chan UnregisterMessage
}

type Message struct {
	Route   string
	Client  *websocket.Conn
	Message []byte
}

type RegisterMessage struct {
	Client *websocket.Conn
	Route  string
}

type UnregisterMessage struct {
	Client *websocket.Conn
	Route  string
}

func NewHub() HubInterface {
	return &Hub{
		clients:    make(map[*websocket.Conn]map[string]bool),
		broadcast:  make(chan Message),
		register:   make(chan RegisterMessage),
		unregister: make(chan UnregisterMessage),
	}
}

func (h *Hub) Register(client *websocket.Conn, route string) {
	h.register <- RegisterMessage{Client: client, Route: route}
}

func (h *Hub) Unregister(client *websocket.Conn, route string) {
	h.unregister <- UnregisterMessage{Client: client, Route: route}
}

func (h *Hub) Broadcast(client *websocket.Conn, route string, message []byte) {
	h.broadcast <- Message{Route: route, Message: message, Client: client}
}

func (h *Hub) Run() {
	for {
		select {
		case register := <-h.register:
			clients, ok := h.clients[register.Client]
			if !ok {
				clients = make(map[string]bool)
				h.clients[register.Client] = clients
			}
			clients[register.Route] = true

		case unregister := <-h.unregister:
			clients, ok := h.clients[unregister.Client]
			if ok {
				delete(clients, unregister.Route)
				if len(clients) == 0 {
					delete(h.clients, unregister.Client)
				}
			}

		case message := <-h.broadcast:
			for client, routes := range h.clients {
				if client != message.Client {
					if routes[message.Route] {
						err := client.WriteMessage(websocket.TextMessage, message.Message)
						if err != nil {
							log.Println("error broadcasting to client ::: ", zap.Error(err))
						}
					}
				}
			}

		}
	}
}
