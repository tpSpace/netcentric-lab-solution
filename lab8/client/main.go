package client

func main() {
	// listen to the broadcast message
	go ListenForBroadcast.ListenForBroadcast()
}
