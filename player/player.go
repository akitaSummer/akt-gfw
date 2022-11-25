package player

import (
	"akt-gfw/define"
)

type Player struct {
	Uid            uint64
	FriendList     []uint64 // 好友
	HandlerParamCh chan define.HandlerParam
	handlers       map[string]Handler
}

func NewPlayer() *Player {
	p := &Player{
		Uid:        0,
		FriendList: make([]uint64, 0),
		handlers:   make(map[string]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}
