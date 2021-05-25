package net

import (
	"log"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */


var HandlerFunc = make(map[int]func(session *Session,data []byte) ([]byte,error))

func init(){
	HandlerFunc[0] = func(session *Session,data []byte) ([]byte, error) {
		log.Printf("Handler Request: %v!",string(data))
		return []byte("handler Success!"),nil
	}
}