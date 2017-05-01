package websocketserver

// Hub: contains all the clients and send broadcast
// of data to them.
type Hub struct {
	Clients				map[*Client]bool	// List of the clients.
	RegistreChannel		chan *Client		// Channel registre a new client.
	UnregistreChannel 	chan *Client		// Channel to unregistre a client.
}

// NewHub: returns a new hub with all the components
// and ready for action.
func NewHub() *Hub {
	return &Hub {
		Clients: 			make(map[*Client]bool),	// List of the clients.
		RegistreChannel: 	make(chan *Client),		// Channel to registre clients.
		UnregistreChannel:	make(chan *Client),		// Channel to unregistre clients.
	}
}

// Run: handles all the actions that affect directly to
// the client (connection, disconnection, client access).
func (h *Hub) Run() {
	// Run a goroutine to read all the messages.
	go func() {
		for {
			select {
			// Add new client to the system.
			case client := <-h.RegistreChannel:
				h.Clients[client] = true
			// Remove a client.
			case client := <-h.UnregistreChannel:
				// Check if the user exist.
				if _, state := h.Clients[client]; state {
					// Delete the user from the user list.
					delete(h.Clients, client)
				}
			}
		}
	}()
}