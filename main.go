package main

import (
	"fmt"
	"main/core"
	"main/netpkg"
	"main/utils"
	"main/wallet"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func main(){
	server := &netpkg.TcpServer{
		Address:  "127.0.0.1",
		Port:     "7777",
		StopChan: make(chan struct{},1),
	}
	go func() {
		server.Listen()
	}()
	keys := utils.LoadWallet("./wallet/sercurt.key")
	fmt.Println(string(keys.GetAddress()))
	ok := utils.IsVaildBitcoinAddress(string(keys.GetAddress()))
	if !ok{
		panic("err")
	}
	core.Mining(keys.GetAddress(),keys.PrivateKey)
	fmt.Printf("amount %.6f\n",wallet.GetWalletBalance(keys.PublicKey))
	bx,_ := wallet.WalletTransfer(keys.PublicKey,keys.PrivateKey,[]byte("1Q1w7NaikzaDgYZKngope6hnMLofok85tj"),2)
	cli := wallet.WalletClient{}
	//
	cli.Listen()
	//
	cli.SendMsg(bx)
	//
	select {}
	//wallet.GetBalance()

	//bitcoinAddress := keys.GetAddress()
	//fmt.Println("TulipCoin Address:", string(bitcoinAddress))
	//fmt.Printf("Verify TulipCoin Address:%v\n", wallet.IsVaildBitcoinAddress(string(bitcoinAddress)))
	//fmt.Printf("%s\n",keys.GetPrivateKey())
	////core.Mining()
	//dax ,_ := keys.PrivateKey.Sign(rand.Reader,[]byte("xxxx"),nil)
	//fmt.Println(dax)
	//xx := wallet.WalletAddressToPublicKeyHash([]byte("15sUz8kH4iKYtMyZbgac3rXjPfbtruPVTj"))
	//fmt.Println(xx)
	//transaction := core.NewTransaction(keys.PublicKey,xx)
	//core.SignTransaction(transaction,keys.PrivateKey)
}