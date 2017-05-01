package websocketserver

import (
	"encoding/json"
	"fmt"
)

// JsonConfiguration: add deseralization behavior.
type JsonConfiguration interface {
	Parse([] byte) error
}

// GeneralConfiugration: handles the main configuration
// file, it contains the SERVER configuration.
type GeneralConfiugration struct {
	Server	ServerConfiguration	`json:"Server"`
}

// Implementing JsonConfiguration interface.
func (gc *GeneralConfiugration) Parse(b []byte) error {
	// Parse and set data into the configuration
	// structure.
	return json.Unmarshal(b, &gc)
}

// ServerConfiguration: handle the SERVER configuration
// it contains the properties needed to boot a nice server.
type ServerConfiguration struct {
	// Hostname used to start the server.
	Host			string	`json:"Hostname"`
	// Port used to start the server.
	Port			int		`json:"Port"`
	// Number of workers to create.
	NumberOfWorkers	int8	`json:"NumberOfWorkers"`
	// Max number of jobs.
	MaxNumberOfJobs	int8	`json:"MaxNumberOfJobs"`
	// Max number of
}

// GetHostname: return the full host name (host and port)
func (sc *ServerConfiguration) GetHostname() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

// Implementing JsonConfiguration interface.
func (sc ServerConfiguration) Parse(b []byte) error {
	// Parse and set data into the configuration
	// structure.
	return json.Unmarshal(b, &sc)
}