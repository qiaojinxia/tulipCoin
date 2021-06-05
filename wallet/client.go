package wallet

import (
	"fmt"
	"main/utils"
	"net"
	"main/netpkg"
)

/**
 * Created by @CaomaoBoy on 2021/5/25.
 *  email:<115882934@qq.com>
 */

type WalletClient struct {
	conn net.Conn
}
func (wc *WalletClient) SendMsg(msg []byte){
	_, err := wc.conn.Write(netpkg.Pack(msg,0))
	if err != nil{
		panic(utils.NetErroWarp(err,""))
	}
}

func(wc *WalletClient) Listen(){
	conn, err := net.Dial("tcp4", "127.0.0.1:7777")
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	wc.conn = conn
}