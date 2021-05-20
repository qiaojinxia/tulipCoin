package wallet

import (
	"bytes"
	"fmt"
	"log"
	"main/core"
	"main/utils"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2021/5/19.
 *  email:<115882934@qq.com>
 */

func GetBalance(){
	txs,err := utils.GetDb().GetAllTransactions()
	if err != nil{
		log.Panic(err)
	}
	transIter := core.NewTransactionIter(txs)
	for transIter.HasNext(){
		fmt.Println(transIter.Next())
	}
}

func GetUtxoSpended() map[string][]core.TxInput{
	utxoSpend :=  make(map[string][]core.TxInput)
	txs,err := utils.GetDb().GetAllTransactions()
	if err != nil{
		log.Panic(err)
	}
	transIter := core.NewTransactionIter(txs)
	for transIter.HasNext(){
		tx := transIter.Next()
		if tx.IsCoinbase() {
			continue
		}
		for _,vin := range tx.Vin{
			utxoSpend[string(tx.ID)] = append(utxoSpend[string(tx.ID)], *vin)
		}

	}
	return utxoSpend
}

func GetUtxoUnSpend(allUseUtxo map[string][]core.TxInput) map[string][]core.TxOutput{
	txs,err := utils.GetDb().GetAllTransactions()
	if err != nil{
		log.Panic(err)
	}
	transIter := core.NewTransactionIter(txs)
	utxoUnSpendTxOut :=  make(map[string][]core.TxOutput)
	for transIter.HasNext(){
		tx := transIter.Next()
		for _,vout := range tx.Vout{
			isuse := false
			if txVins,ok := allUseUtxo[string(tx.ID)];ok{
				for _,txVin := range txVins{
					if txVin.Sequence == vout.No {
						isuse = true
					}
				}

			}
			if !isuse {
				utxoUnSpendTxOut[string(tx.ID)] = append(utxoUnSpendTxOut[string(tx.ID)], *vout)
			}
		}
	}
	return utxoUnSpendTxOut
}


func GetWalletAddressBalance(walletAddress []byte,allOutPut map[string][]core.TxOutput) float64{
	charge :=  0.0
	walletPublicKeyHash := utils.WalletAddressToPublicKeyHash(walletAddress)
	for _,txOuts := range allOutPut{
		for _, txOut := range txOuts{
			publicHash := []byte(strings.Split(txOut.ScriptPubKey," ")[2])
			if bytes.Equal(publicHash,walletPublicKeyHash) {
				charge += txOut.Value
			}
		}
	}
	return charge
}