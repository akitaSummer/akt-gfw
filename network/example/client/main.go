package main

import (
	"akt-gfw/network"
)

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
