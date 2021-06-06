package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"main/config"
	"main/utils"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

type Transaction struct {
	TxID      []byte
	Vin       []*TxInput
	Vout      []*TxOutput
	TimeStamp int64

}
func(tx *Transaction) SetID(){
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil{
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TxID = hash[:]
}


func(tx *Transaction) IsCoinbase() bool{
	if  len(tx.Vin) ==1 && len(tx.Vin[0].PrevTxHash) == 0 && tx.Vin[0].Vout == -1{
		return true
	}
	return false
}

type TxInput struct {
	TxID       []byte
	Vout       int
	PrevTxHash []byte
	ScriptSig  string //签名 和 公钥
}


func(ti *TxInput) String() string{
	return fmt.Sprintf("Vout:%d , PrevTxHash:%x , ScriptSig:%s \n",ti.Vout,ti.PrevTxHash,ti.ScriptSig)
}

//<PubK(B)> OP_DUP OP_HASH160 <PubKHash(B)> OP_EQUALVERIFY OP_CHECKSIG
type TxOutput struct {
	No int
	Value float64
	ScriptPubKey string
}


func(ti *TxOutput) String() string{
	return fmt.Sprintf("VoutNo:%d , TransferVlaue:%.7f , ScriptPubKey:%s \n",ti.No,ti.Value,ti.ScriptPubKey)
}


func NewCoinbase(toAddress []byte,data string) *Transaction{
	if data == ""{
		data = fmt.Sprintf("Reward to '%s'",toAddress)
	}
	publicKeyHash:= fmt.Sprintf("OP_DUP OP_HASH160 %x OP_EQUALVERIFY OP_CHECKSIG ",utils.WalletAddressToPublicKeyHash(toAddress))
	tx := &Transaction{
		Vin:  []*TxInput{{PrevTxHash:[]byte{} ,ScriptSig: data, Vout: -1}},
		Vout: []*TxOutput{{Value:config.RewardCoinn,ScriptPubKey: publicKeyHash,No:1}},
		TimeStamp: time.Now().UnixNano() /1e6,
	}
	tx.SetID()
	return tx
}

func NewTransaction(publicKey []byte,toWalletAddress []byte,txOuts map[string]TxOutput,amount float64) (*Transaction,error){
	txInputs := make([]*TxInput,0,len(txOuts))
	txOutputs := make([]*TxOutput,0)
	sum := 0.0

	for txID,txOut := range txOuts{
		trans ,_ := utils.GetDb().GetTransactionByTxID([]byte(txID))
		if trans == nil{
			panic(utils.BusinessErrorWarp(errors.New(""),"Transaction Can't Find!"))

		}
		sum += txOut.Value
		txInputs = append(txInputs, &TxInput{
			Vout:       txOut.No,
			PrevTxHash: []byte(txID),
			ScriptSig:  string(publicKey),
		})

	}
	fmt.Println("test",fmt.Sprintf("%x",utils.GeneratePublicKeyHash(publicKey)))
	transferOut := &TxOutput{
		No:           len(txOutputs),
		Value:        amount,
		ScriptPubKey: utils.GenerateLockScript(toWalletAddress),
	}
	txOutputs = append(txOutputs, transferOut)
	//surplus change transfer to self address
	if sum > amount {
		surplus := sum - amount
		sTo := &TxOutput{
			No:           len(txOutputs),
			Value:        surplus,
			ScriptPubKey: utils.GenerateLockScriptByPublicHash(publicKey),
		}
		txOutputs = append(txOutputs, sTo)
	}else if sum < amount{
		return nil,errors.New("Not Enought UTXO Can Transfer!")
	}
	tx := &Transaction{
		TxID: nil,
		Vin:  txInputs,
		Vout: txOutputs,
	}
	return tx,nil
}

func SignTransaction(transaction *Transaction,privateKey *ecdsa.PrivateKey){
	if transaction.IsCoinbase() {
		return
	}
	cpTransaction := CopyTX(transaction)
	for index := range cpTransaction.Vin{
		publicKey := cpTransaction.Vin[index].ScriptSig
		transaction.SetID()
		sindata := transaction.TxID
		r,s,err := ecdsa.Sign(rand.Reader,privateKey,sindata)
		if err != nil{
			log.Panic(err)
		}
		signature := append(r.Bytes(),s.Bytes()...)
		transaction.Vin[index].TxID = []byte(fmt.Sprintf("%x",sindata))
		transaction.Vin[index].ScriptSig = fmt.Sprintf("%x %x",signature,publicKey)
	}
	transaction.SetID()
}

func CopyTX(tarnsaction *Transaction) Transaction{
	newTrans := Transaction{}
	vin := make([]*TxInput,0,len(tarnsaction.Vin))
	for _,vi := range tarnsaction.Vin{
		txin := &TxInput{
			Vout:       vi.Vout,
			PrevTxHash: vi.PrevTxHash,
			ScriptSig:  vi.ScriptSig,
		}
		vin = append(vin, txin)
	}
	vot:= make([]*TxOutput,0,len(tarnsaction.Vout))
	for _,vo := range tarnsaction.Vout{
		txin := &TxOutput{
			No:           vo.No,
			Value:        vo.Value,
			ScriptPubKey: vo.ScriptPubKey,
		}
		vot = append(vot, txin)
	}
	newTrans.Vout = vot
	newTrans.Vin = vin
	return newTrans
}