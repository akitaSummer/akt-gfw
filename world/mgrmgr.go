package world

import "akt-gfw/manager"

var MM *MgrMgr

type MgrMgr struct {
	Pm manager.PlayerMgr
}

func NewMgrMgr() *MgrMgr {
	return &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
}
