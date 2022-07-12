package cli

type Hub struct {
	clients map[*Cli]bool

	broadcast chan []byte

	register chan *Cli

	unregister chan *Cli
}

func InitHub() *Hub {
	return &Hub{
		clients:    make(map[*Cli]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Cli),
		unregister: make(chan *Cli),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
