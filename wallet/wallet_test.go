package wallet

import (
	"fmt"
	"main/utils"
	"os"
	"testing"
)

/**
 * Created by @CaomaoBoy on 2021/5/19.
 *  email:<115882934@qq.com>
 */

func TestGetUtxoSpended(t *testing.T) {
	spendUtxo := GetUtxoSpended()
	fmt.Printf("%v",spendUtxo)
}

func TestGetUtxoUnSpend(t *testing.T) {
	spendUtxo := GetUtxoSpended()
	unSpendUtxo := GetUtxoUnSpend(spendUtxo)
	fmt.Printf("%v",unSpendUtxo)
}

func TestGetWalletAddressBalance(t *testing.T) {
	spendUtxo := GetUtxoSpended()
	unSpendUtxo := GetUtxoUnSpend(spendUtxo)
	GetWalletAddressBalance([]byte("1Q1w7NaikzaDgYZKngope6hnMLofok85tj"),unSpendUtxo)
	//fmt.Printf("Surplus Amount %.6f", val)
}

func TestWalletTransfer(t *testing.T) {
	//WalletTransfer()
}

func TestStorWallet(t *testing.T) {
	file,err := os.OpenFile("./sercurt.key",os.O_WRONLY|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err!= nil{
		panic(err)
	}
	defer file.Close()
	keys := utils.GetBitcoinKeys()
	fmt.Printf("%s\n",keys.GetAddress())
	privateKey,publicKey := keys.KeysString()
	key := fmt.Sprintf("%s;%s",privateKey,publicKey)
	file.Write([]byte(key))
	xs := utils.BitcoinKeys{}
	xs.Init(privateKey,publicKey)
	fmt.Printf("%s\n",xs.GetAddress())
}