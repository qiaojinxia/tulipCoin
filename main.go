package main

import (
	"main/core"
	"main/wallet"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func main(){
	keys := wallet.GetBitcoinKeys()
	core.Mining(keys.GetAddress(),keys.PrivateKey)
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