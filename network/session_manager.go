package network

import "sync"

type SessionMgr struct {
	Sessions map[int64]*Session
	Counter  int64 //计数器
	Mutex    sync.Mutex
	Pid      int64
}

var (
	SessionMgrInstance SessionMgr
	oneInitSessionMgr  sync.Once
)

func init() {
	oneInitSessionMgr.Do(func() {
		SessionMgrInstance = SessionMgr{
			Sessions: make(map[int64]*Session),
			Counter:  0,
			Mutex:    sync.Mutex{},
		}
	})
}

func (sm *SessionMgr) AddSession(s *Session) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()
	if val := sm.Sessions[s.UId]; val != nil {
		if val.IsClose {
			sm.Sessions[s.UId] = s
		} else {
			return
		}
	}
}

func (sm *SessionMgr) DelSession(UId int64) {
	delete(sm.Sessions, UId)
}
