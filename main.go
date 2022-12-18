package main

import (
	"akt-gfw/world"

	"github.com/phuhao00/sugar"
)

func main() {
	world.MM = world.NewMgrMgr()

	go world.MM.Pm.Run()

	sugar.WaitSignal(world.MM.OnSystemSignal)
}
