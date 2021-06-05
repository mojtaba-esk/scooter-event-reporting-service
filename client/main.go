package main

import (
	"fmt"
	"os"
	"strconv"
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

	numOfClients, err := strconv.Atoi(ENV.NUM_OF_CLIENTS)
	if err != nil {
		numOfClients = 3
	}
	clients, _ := GetRandomClients(numOfClients)

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
	NUM_OF_CLIENTS   string
}

func init() {

	ENV.SCOOTIN_API_PATH = os.Getenv("SCOOTIN_API_PATH")
	ENV.STATIC_API_KEY = os.Getenv("STATIC_API_KEY")
	ENV.NUM_OF_CLIENTS = os.Getenv("NUM_OF_CLIENTS")

}
