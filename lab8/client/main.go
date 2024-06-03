package client

import (
	"../ListenForBroadcast.go"
)

func main() {
	// listen to the broadcast message
	go ListenForBroadcast.ListenForBroadcast()
}
