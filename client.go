package websocketserver

import(
	"github.com/gorilla/websocket"
)

// Client: handles the connection.
type Client struct {
	Hub			*Hub			// A reference to the hub.
	Connection	*websocket.Conn	// The websocket connection.
	JobQueue	chan Job		// Assigned queue.
}

// ReadData: waits for a new message and send
// it to the
func (c *Client) ReadData() {
	// When the client dead.
	defer func() {
		c.Hub.UnregistreChannel <- c
		c.Connection.Close()
	}()
	// Keep alive the connection.
	for {
		// Wait for a message.
		_, messageBuffer, err := c.Connection.ReadMessage()
		// The connection has a problem.
		// It has to stop the connection.
		if err != nil { break }

		newJob := Job {
			Data:	messageBuffer,
			Client:	c,
		}

		c.JobQueue <- newJob
	}
}