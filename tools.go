package websocketserver

import "fmt"

// ShowRunMessage: show the starting server message.
func ShowRunMessage(conf ServerConfiguration) {
	fmt.Println("[*] Starting web socket server")
	fmt.Println("\t[Host]:        ", conf.Host)
	fmt.Println("\t[Port]:        ", conf.Port)
	fmt.Println("\t[Num workers]: ", conf.NumberOfWorkers)
	fmt.Println("\t[Num jobs]:    ", conf.MaxNumberOfJobs)
}