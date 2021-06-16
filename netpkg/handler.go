package netpkg

import (
	"fmt"
	"main/core"
	"main/dto"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */



const (
	ReceiveTransaction int = 0
	ConnectionVerify int = 1
	GetHeaders int = 2
	HeartBeat int = 3
	NodeOnline int = 4
)

var ClientHandlerFunc = make(map[int]func(session *Session,data []byte) ([]byte,error))
var ServerHandlerFunc = make(map[int]func(session *Session,data []byte) ([]byte,error))

func init(){

	ClientHandlerFunc[ReceiveTransaction] = func(session *Session,data []byte) ([]byte, error) {
		transaction := dto.ConvertTransactionBytes(data)
		ctp := core.GetCtxPool()
		tmpt := &core.CTxMemPoolEntry{CTransactionRef: transaction}
		ctp.AddTxToPool(tmpt)
		return []byte("handler Success!"),nil
	}

	ClientHandlerFunc[ConnectionVerify] = func(session *Session,data []byte) ([]byte, error) {
		transaction := dto.ConvertTransactionBytes(data)
		ctp := core.GetCtxPool()
		tmpt := &core.CTxMemPoolEntry{CTransactionRef: transaction}
		ctp.AddTxToPool(tmpt)
		return []byte("handler Success!"),nil
	}

	ClientHandlerFunc[NodeOnline] = func(session *Session,data []byte) ([]byte, error) {
		//Sync Block
		blockHearders,err := utils.GetDb().IterAllBlock()
		if err != nil{
			panic(utils.BusinessErrorWarp(err,""))
		}
		fmt.Println(blockHearders)
		return []byte("handler Success!"),nil
	}


}

