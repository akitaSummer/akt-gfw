package network

import "sync"

type SessionMgr struct {
	Sessions map[uint64]*Session
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
			Sessions: make(map[uint64]*Session),
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

func (sm *SessionMgr) DelSession(UId uint64) {
	delete(sm.Sessions, UId)
}
