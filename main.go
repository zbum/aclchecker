package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	toolType := flag.String("type", "client", "The type of role to assign to this process. It can be either 'client' or 'server'.")
	protocol := flag.String("protocol", "tcp", "The type of role to assign to this process. It can be either 'client' or 'server'.")
	flag.Parse()

	switch *toolType {
	case "client":
		config := LoadClientConfig()
		StartClient(*protocol, config)
	case "server":
		config := LoadServerConfig()
		StartServer(*protocol, config)

		// Wait for a signal to terminate the server
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		<-done
	default:
	}

}
