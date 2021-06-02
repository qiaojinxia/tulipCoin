package core

import (
	"crypto/ecdsa"
	"main/config"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */

type Block struct {
	*Header
	*Body `json:"-"`
}

type Header struct {
	Index int64 `json:"index"`
	PreviousHash []byte `json:"previous_hash"`
	Memorials  []*Memorial `json:"memorial"`
	TimeStamp int64 `json:"time_stamp"`
	MRoot  []byte `json:"root"`
	Nonce int64 `json:"nonce"`
	Version string `json:"version"`
	Diffcult int32 `json:"diffcult"`
	Hash []byte `json:"hash"`
}

type Body struct {
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(index int,prevHash []byte,toAddress []byte,memo string,privateKey *ecdsa.PrivateKey) *Block{
	baseCoin := NewCoinbase(toAddress,memo)
	block := &Block{
		&Header{
			Index:        int64(index),
			PreviousHash: prevHash,
			TimeStamp:    time.Now().UnixNano(),
			Version:      config.Version,
			Diffcult:	  config.TargetBits,
		},
		&Body{
			Transactions:  []*Transaction{
				baseCoin,
			},
		},

	}
	//TO get transaction from pool
	for i:=0;i<10;i++{
		if trans := GetCtxPool().PopTx();trans == nil{
			break
		}else{
			block.Transactions = append(block.Transactions, trans.CTransactionRef)
		}
	}
	transesID := make([][]byte,0,len(block.Transactions))
	for _,tx := range block.Transactions{
		SignTransaction(tx,privateKey)
		transesID = append(transesID,tx.TxID)
	}
	mRoot := NewMerkleTree(transesID)
	block.MRoot = mRoot.merkleRoot
	return block
}

//创造创世区块
func CreateGenesisBlock() *Block{
	return &Block{
		&Header{
			Index:        1,
			PreviousHash: []byte{},
			TimeStamp:    0,
			MRoot: []byte(""),
			Hash:         []byte(""),
			Version:      config.Version,
		},&Body{},
	}
}

