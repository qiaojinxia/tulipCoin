package main

import (
	"crypto/rand"
	"fmt"
	"main/wallet"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func main(){
	keys := wallet.GetBitcoinKeys()
	bitcoinAddress := keys.GetAddress()
	fmt.Println("TulipCoin Address:", string(bitcoinAddress))
	fmt.Printf("Verify TulipCoin Address:%v\n", wallet.IsVaildBitcoinAddress(string(bitcoinAddress)))
	fmt.Printf("%s\n",keys.GetPrivateKey())
	//core.Mining()
	dax ,_ := keys.PrivateKey.Sign(rand.Reader,[]byte("xxxx"),nil)
	fmt.Println(dax)

}