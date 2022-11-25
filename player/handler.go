package player

import (
	"akt-gfw/chat"
	"akt-gfw/function"
	"fmt"
)

type Handler func(interface{})

func (p *Player) AddFirend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFirend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg)
	fmt.Println(chatMsg)
}
