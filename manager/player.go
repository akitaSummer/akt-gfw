package manager

import "akt-gfw/player"

// 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player
	addPch  chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.Uid] = p
	go p.Run()
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPch:
			pm.Add(p)
		}
	}
}
