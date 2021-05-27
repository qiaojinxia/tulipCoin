package netpkg

import (
	"main/core"
	"main/dto"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */


var HandlerFunc = make(map[int]func(session *Session,data []byte) ([]byte,error))

func init(){
	HandlerFunc[0] = func(session *Session,data []byte) ([]byte, error) {
		transaction := dto.ConvertTransactionBytes(data)
		ctp := core.GetCtxPool()
		tmpt := &core.CTxMemPoolEntry{CTransactionRef: transaction}
		ctp.AddTxToPool(tmpt)
		return []byte("handler Success!"),nil
	}
}