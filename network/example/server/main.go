package main

import "akt-gfw/network"

func main() {
	server := network.NewServer(":8023", "tcp6")
	server.Run()
	select {}
}
