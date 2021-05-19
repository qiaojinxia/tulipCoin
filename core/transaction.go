package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"main/config"
	"main/wallet"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

type Transaction struct {
	ID []byte
	Vin []*TxInput
	Vout []*TxOutput

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
	tx.ID = hash[:]
}


func(tx *Transaction) IsCoinbase() bool{
	if  len(tx.Vin) ==1 && len(tx.Vin[0].PrevTxHash) == 1 && tx.Vin[0].Sequence == -1{
		return true
	}
	return false
}

type TxInput struct {
	Sequence int
	PrevTxHash []byte
	ScriptSig string //签名 和 公钥
}

func(ti *TxInput) String() string{
	return fmt.Sprintf("Sequence:%d , PrevTxHash:%x , ScriptSig:%s \n",ti.Sequence,ti.PrevTxHash,ti.ScriptSig)
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
	publicKeyHash:= fmt.Sprintf("OP_DUP OP_HASH160 <%x> OP_EQUALVERIFY OP_CHECKSIG ", wallet.WalletAddressToPublicKeyHash(toAddress))
	tx := &Transaction{
		Vin:  []*TxInput{{PrevTxHash:[]byte{} ,ScriptSig: data,Sequence: -1}},
		Vout: []*TxOutput{{Value:config.RewardCoinn,ScriptPubKey: publicKeyHash,No:1}},
	}
	tx.SetID()
	return tx
}

func NewTransaction(publicKey []byte,publicKeyHash []byte) *Transaction{
	tx := &Transaction{
		ID:   nil,
		Vin:  []*TxInput{{
			Sequence:      0,
			PrevTxHash:    []byte("asdasdasdasd12312"),
			ScriptSig: fmt.Sprintf("%x",publicKey),
		},
		},
		Vout: []*TxOutput{
			{No:1,
			Value:60.0,
			ScriptPubKey:fmt.Sprintf("OP_DUP OP_HASH160 <%x> OP_EQUALVERIFY OP_CHECKSIG ",publicKeyHash)},
		},
	}
	return tx
}

func SignTransaction(transaction *Transaction,privateKey *ecdsa.PrivateKey){
	if transaction.IsCoinbase() {
		return
	}
	cpTransaction := CopyTX(transaction)
	for index := range cpTransaction.Vin{
		bak := cpTransaction.Vin[index].ScriptSig
		transaction.SetID()
		sindata := transaction.ID
		r,s,err := ecdsa.Sign(rand.Reader,privateKey,sindata)
		if err != nil{
			log.Panic(err)
		}
		signature := append(r.Bytes(),s.Bytes()...)
		transaction.Vin[index].ScriptSig = fmt.Sprintf("<%x> <%x>",signature,bak)
	}
	transaction.SetID()
}

func CopyTX(tarnsaction *Transaction) Transaction{
	newTrans := Transaction{}
	vin := make([]*TxInput,0,len(tarnsaction.Vin))
	for _,vi := range tarnsaction.Vin{
		txin := &TxInput{
			Sequence:      vi.Sequence,
			PrevTxHash:    vi.PrevTxHash,
			ScriptSig: 	   vi.ScriptSig,
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