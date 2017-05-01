package websocketserver

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Handles the requests made by the user.
type RequestHandler func(job Job)

// Sever: handles the request made by
// the clients.
type Server struct {
	Configuration 	ServerConfiguration
	Channel       	string
	Hub				*Hub
	JobQueue		chan Job
	Handler			RequestHandler
}

// NewServer: creates and return a new server
// instance.
func NewServer(conf ServerConfiguration) Server {
	return Server{
		// Current configuration.
		Configuration: 	conf,
		Hub:			NewHub(),
		JobQueue:		make(chan Job, conf.MaxNumberOfJobs),
	}
}

// Run: start the server.
func Run(wsocketServer Server) {
	// Show the init server message.
	ShowRunMessage(wsocketServer.Configuration)
	// Run the HUB
	wsocketServer.Hub.Run()
	// Dispatch all the workers.
	dispatchWorkers(
		&wsocketServer,
		wsocketServer.JobQueue,
		wsocketServer.Configuration.NumberOfWorkers,
	)
	// Listen in all the channels.
	http.HandleFunc(wsocketServer.Channel, func(w http.ResponseWriter, r *http.Request) {
		var newConnection, _ = upgrader.Upgrade(w, r, nil)
		// Shit happen, just avoid code.
		if newConnection == nil { return }
		// Create a new client.
		newClient := &Client {
			Hub:			wsocketServer.Hub,
			Connection:		newConnection,
			JobQueue:		wsocketServer.JobQueue,
		}
		// Registre the new client in the hub.
		wsocketServer.Hub.RegistreChannel <- newClient
		// Listen to the client.
		go newClient.ReadData()
	})
	// Start the http server.
	http.ListenAndServe(wsocketServer.Configuration.GetHostname(), nil)
}

// dispatchWorkers: dispatches a specific number of workers.
func dispatchWorkers(server *Server, jobQueue chan Job, numberOfWorkers int8) {
	var i int8 = 0
	for i = 1; i <= numberOfWorkers; i++ {
		go func() {
			for {
				select {
				case job := <- jobQueue:
					server.Handler(job)
				}
			}
		}()
	}
}
