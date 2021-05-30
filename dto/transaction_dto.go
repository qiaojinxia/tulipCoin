package dto

import (
	"github.com/golang/protobuf/proto"
	"main/core"
	"main/protomsg"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/25.
 *  email:<115882934@qq.com>
 */

func PConvertTransactionBytes(transaction *core.Transaction) ([]byte,error){
	tx := &protomsg.Transaction{
		ID:                   transaction.TxID,
		Vin:                  make([]*protomsg.TxInput,0),
		Vount:                make([]*protomsg.TxOutput,0),
	}
	for _,vi := range transaction.Vin{
		p_vin := &protomsg.TxInput{
			TxID: 				  vi.TxID,
			Vout:            	  int64(vi.Vout),
			PrevTxHash:           vi.PrevTxHash,
			ScriptSig:            vi.ScriptSig,

		}
		tx.Vin = append(tx.Vin, p_vin)
	}
	for _,vot := range transaction.Vout{
		p_vot := &protomsg.TxOutput{
			No:                   int64(vot.No),
			Value:                float32(vot.Value),
			ScriptPubKey:         vot.ScriptPubKey,
		}
		tx.Vount = append(tx.Vount, p_vot)
	}

	return proto.Marshal(tx)
}

func ConvertTransactionBytes(pTranssactions []byte) *core.Transaction{
	pTrans := &protomsg.Transaction{}
	err := proto.Unmarshal(pTranssactions,pTrans)
	if err != nil{
		utils.MarshalErrorWarp(err.Error())
	}
	vins := make([]*core.TxInput,0,len(pTrans.Vin))
	vouts := make([]*core.TxOutput,0,len(pTrans.Vount))
	for _,pVin := range pTrans.Vin{
		vin := &core.TxInput{
			Vout:       int(pVin.Vout),
			TxID: pVin.TxID,
			PrevTxHash: pVin.PrevTxHash,
			ScriptSig:  pVin.ScriptSig,
		}
		vins = append(vins, vin)
	}

	for _,pVot := range pTrans.Vount{
		vout := &core.TxOutput{
			No:           int(pVot.No),
			Value:       float64(pVot.Value),
			ScriptPubKey: pVot.ScriptPubKey,
		}
		vouts = append(vouts, vout)
	}
	trans := &core.Transaction{
		TxID: pTrans.ID,
		Vin:  vins,
		Vout: vouts,
	}
	return trans
}