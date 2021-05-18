package core

import (
	"main/config"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */

type Block struct {
	Index int64 `json:"index"`
	PreviousHash []byte `json:"previous_hash"`
	Memorials  []*Memorial `json:"memorial"`
	TimeStamp int64 `json:"time_stamp"`
	Transactions []*Transaction `json:"transactions"`
	MRoot  []byte `json:"root"`
	Nonce int64 `json:"nonce"`
	Version string `json:"version"`
	Hash []byte `json:"hash"`
}

func NewBlock(index int,prevHash []byte,toAddress []byte,memo string) *Block{
	baseCoin := NewCoinbase(toAddress,memo)
	return &Block{
		Index:        int64(index),
		PreviousHash: prevHash,
		TimeStamp:    time.Now().UnixNano(),
		Transactions:  []*Transaction{
			baseCoin,
		},
		Version:      config.Version,
	}
}

//创造创世区块
func CreateGenesisBlock() *Block{
	return &Block{
		Index:        1,
		PreviousHash: []byte{},
		TimeStamp:    0,
		Transactions: make([]*Transaction,0),
		Hash:         []byte(""),
		Version:      config.Version,
	}
}

