package wallet

import (
	"fmt"
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
	val := GetWalletAddressBalance([]byte("1Q1w7NaikzaDgYZKngope6hnMLofok85tj"),unSpendUtxo)
	fmt.Println(val)
}