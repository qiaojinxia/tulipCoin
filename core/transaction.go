package core

import (
	"fmt"
	"main/config"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

type Transaction struct {
	ID []byte
	Vin TxInput
	Vout TxOutput

}
func(tx *Transaction) SetID(){

}

type TxInput struct {
	Txid []byte
	Vout int
	ScriptSig string
}

type TxOutput struct {
	Value int
	ScriptPubKey string
}

func NewCoinbaseTX(to,data string) *Transaction{
	if data == ""{
		data = fmt.Sprintf("Reward to '%s'",to)
	}
	txin := TxInput{Txid:[]byte{}, Vout:config.RewardCoinn ,ScriptSig: data, }
	txout := TxOutput{Value:1,ScriptPubKey: to,}
	tx := &Transaction{
		ID:   nil,
		Vin:  txin,
		Vout: txout,
	}
	tx.SetID()
	return tx
}