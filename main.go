package main

import (
	"fmt"
	"log"
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


	//aa := fmt.Sprintf("%x",keys.PublicKey)
	xx := utils.GeneratePublicKeyHash([]byte("304402200fec574bdab69adf378b488671afade0cee88ddac74f27caef5c4183f43265190220330e14b202e3f1134dee1e37a90af3bb9970ea2302223624d71b774aa5c4c91801"))
	fmt.Printf("xxx %x",xx)

	server := &netpkg.TcpServer{
		Address:  "127.0.0.1",
		Port:     "7777",
		StopChan: make(chan struct{},1),
	}
	go func() {
		server.Listen()
	}()
	utils.Try(
		func() {
			keys := utils.LoadWallet("./wallet/sercurt.key")
			fmt.Println(string(keys.GetAddress()))
			ok := utils.IsVaildBitcoinAddress(string(keys.GetAddress()))
			if !ok{
				panic("err")
			}
			core.Mining(keys.GetAddress(),keys.PrivateKey)
			fmt.Printf("amount %.6f\n",wallet.GetWalletBalance(keys.PublicKey))
			fmt.Printf("Public Key %x\n",keys.PublicKey)
			bx,err := wallet.WalletTransfer(keys.PublicKey,keys.PrivateKey,[]byte("1Q1w7NaikzaDgYZKngope6hnMLofok85tj"),2)
			if err != nil{
				panic(utils.ConverErrorWarp(err.Error()))
			}
			cli := wallet.WalletClient{}
			//
			cli.Listen()
			//
			cli.SendMsg(bx)
			//
			select {}
		}).CatchAll(func(err error) {
		log.Panicf("Catch Error : %s!" ,utils.ConverToJsonInfo(err))
	})

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