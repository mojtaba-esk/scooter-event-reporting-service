package main

import (
	"fmt"
	"os"
	"time"
)

/*-------------*/

func main() {

	//Wait for the server to get ready
	for !IsServerReady() {
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("\nServer is ready")

	/*------*/

	clients, _ := GetRandomClients(3)

	for _, client := range clients {
		fmt.Printf("\nStarting Client: %15s\t UUID: %s", client.Name, client.UUID)
		go StartClient(client)
	}

	/*------*/

	//Keep the main routine blocking
	var c chan struct{}
	<-c
}

/*-------------*/

var ENV struct {
	SCOOTIN_API_PATH string
	STATIC_API_KEY   string
}

func init() {

	ENV.SCOOTIN_API_PATH = os.Getenv("SCOOTIN_API_PATH")
	ENV.STATIC_API_KEY = os.Getenv("STATIC_API_KEY")

}
