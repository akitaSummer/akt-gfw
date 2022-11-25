package player

import (
	"akt-gfw/chat"
	"akt-gfw/function"
)

type Player struct {
	Uid        uint64
	FriendList []uint64 // 好友
	chChat     chan chat.Msg
}

func NewPlayer() *Player {
	return &Player{
		Uid:        0,
		FriendList: nil,
	}
}

func (p *Player) AddFirend(fId uint64) {
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFirend(fId uint64) {
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) Run() {
	for {
		select {
		case chatMsg := <-p.chChat:
			p.ResolveChatMsg(chatMsg)
		}
	}
}

func (p *Player) ResolveChatMsg(chatMsg chat.Msg) {}
