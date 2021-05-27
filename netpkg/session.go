package netpkg

import (
	"main/utils"
	"net"
)

/**
 * Created by @CaomaoBoy on 2021/5/22.
 *  email:<115882934@qq.com>
 */

var _SessionManger = SessionManager{Sessions: make(map[int64]*Session)}

type SessionManager struct {
	Sessions map[int64]*Session
}
func(sm *SessionManager) Broadcast(msg []byte){
	for _,v := range sm.Sessions{
		_,err := v.conn.Write(msg)
		if err != nil{
			panic(utils.NetErroWarp(err.Error()))
		}
	}
}

func(sm *SessionManager) AddSession(session *Session){
	if _,ok := sm.Sessions[session.ID];ok{
		return
	}
	sm.Sessions[session.ID] = session
}

func(sm *SessionManager) RemoveSession(ID int64) bool{
	if _,ok := sm.Sessions[ID];ok{
		delete(sm.Sessions, ID)
		return ok
	}
	return false
}

//User Session
type Session struct {
	ID int64
	conn net.Conn
	MsgChan chan []byte
}

func(se *Session) AsyncMsg(){

}
