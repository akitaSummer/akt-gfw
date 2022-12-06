package manager

import "akt-gfw/player"

// 维护在线玩家
type PlayerMgr struct {
	players map[uint64]*player.Player
	addPch  chan player.Player
}

func (pm *PlayerMgr) Add(p *player.Player) {
	pm.players[p.UId] = p
	go p.Run()
}

//Del ...
func (pm *PlayerMgr) Del(p player.Player) {
	delete(pm.players, p.UId)
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPch:
			pm.Add(&p)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
