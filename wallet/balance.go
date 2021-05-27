package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/core"
	"main/dto"
	"main/utils"
	"sort"
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


func GetWalletAddressBalance(walletPublicKeyHash []byte,allUnUsedOutPut map[string][]core.TxOutput) map[string]core.TxOutput{
	ownUnusedUtxo := make(map[string]core.TxOutput)
	for txID,txOuts := range allUnUsedOutPut{
		for _, txOut := range txOuts{
			publicHash := strings.Split(txOut.ScriptPubKey," ")[2]
			if publicHash == fmt.Sprintf("%x",walletPublicKeyHash) {
				ownUnusedUtxo[txID] = txOut
			}
		}
	}
	return ownUnusedUtxo
}

func WalletTransfer(publicKey []byte,privateKey *ecdsa.PrivateKey,toWalletAddress []byte,amount float64) ([]byte,error){
	if !utils.IsVaildBitcoinAddress(string(toWalletAddress)){
		log.Panic("Invalid Wallet!")
	}
	publicKeyHash := utils.GeneratePublicKeyHash(publicKey)
	// 1.Find all the small charge
	spendUtxo := GetUtxoSpended()
	alUnSpendUtxo := GetUtxoUnSpend(spendUtxo)
	utxoOutPut := GetWalletAddressBalance(publicKeyHash,alUnSpendUtxo)
	//2.Find the right amount of change to transfer
	sum := 0.0
	type TempSort struct {
		Key string
		Value float64
	}
	ts := make([]*TempSort,0,len(utxoOutPut))
	for key,value := range utxoOutPut{
		ts = append(ts, &TempSort{
			Key:   key,
			Value: value.Value,
		})

	}
	sort.SliceStable(ts, func(i, j int) bool {
		return ts[i].Value < ts[i].Value
	})
	transferUtxo := make(map[string]core.TxOutput,0)
	for _,txOut := range ts{
		transferUtxo[txOut.Key] = utxoOutPut[txOut.Key]
		sum += txOut.Value
		if sum >= amount{
			break
		}
	}
	// 3.Generate Transaction
	tx,err := core.NewTransaction(publicKey,toWalletAddress,transferUtxo,amount)
	if err != nil{
		log.Panic(err)
	}
	core.SignTransaction(tx,privateKey)
	// 4.Add to UTXO Pool
	return  dto.PConvertTransactionBytes(tx)
}


