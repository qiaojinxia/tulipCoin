package core

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"main/config"
	"main/utils"
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

func GetTargetBit() int32{
	lastBlockIndex,err:= utils.GetDb().GetBlockHeight()
	if lastBlockIndex == 0{
		return config.TargetBits
	}
	bLastBlock,err := utils.GetDb().GetBlock(int64(lastBlockIndex))
	if err != nil{
		panic(utils.DataBaseErrorWarp(err,""))
	}
	lastBlock := &Block{}
	err = json.Unmarshal(bLastBlock,lastBlock)
	if err != nil{
		panic(utils.DataBaseErrorWarp(err,""))
	}
	if lastBlockIndex % config.NInterval != 0{
		return lastBlock.Diffcult
	}
	firstBlock := &Block{}
	firstIndex := lastBlock.Index - config.NInterval
	bFirstBlock,err := utils.GetDb().GetBlock(firstIndex)
	err = json.Unmarshal(bFirstBlock,firstBlock)

	min := lastBlock.Diffcult / 2
	max := lastBlock.Diffcult * 2

	realSpendTime := (lastBlock.TimeStamp - firstBlock.TimeStamp) / 1e3
	TargetBits := lastBlock.Diffcult * int32(realSpendTime/config.NTargetTimespan)
	if max > TargetBits{
		return max
	}else if TargetBits < min{
		return min
	}
	return TargetBits
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
			TimeStamp:    time.Now().UnixNano()/ 1e6,
			Version:      config.Version,
			Diffcult:	  GetTargetBit(),
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
	block.MRoot = []byte(fmt.Sprintf("%x",mRoot.merkleRoot))
	return block
}

//创造创世区块
func CreateGenesisBlock() *Block{
	return &Block{
		&Header{
			Index:        0,
			PreviousHash: []byte{},
			TimeStamp:    time.Now().UnixNano()/ 1e6,
			MRoot: []byte(""),
			Hash:         []byte(""),
			Version:      config.Version,
		},&Body{},
	}
}

