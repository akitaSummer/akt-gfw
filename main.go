package main

import "akt-gfw/world"

func main() {
	world.MM = world.NewMgrMgr()

	world.MM.Pm.Run()
}
