package player

func (p *Player) HandlerRegister() {
	p.handlers["add_friend"] = p.AddFirend
	p.handlers["del_friend"] = p.DelFirend
	p.handlers["resolve_chat_msg"] = p.ResolveChatMsg
}
