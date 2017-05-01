package websocketserver

// Job: contains a specific task made by the user.
type Job struct {
	Data 	[]byte	// The buffer (chunk of data provide by the user).
	Client 	*Client	// The client.
}