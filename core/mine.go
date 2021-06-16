package core

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"main/config"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func Mining(walletAddress []byte,privateKey *ecdsa.PrivateKey){
	lastIndex,preHash := _BlockChain.GetLastBlockIndex()
	block := NewBlock(lastIndex + 1,preHash,walletAddress,"Begin Mine Block,Good Luck!",privateKey)
	pow := NewProofOfWork(block)
	nonce,blockHash := pow.Run()
	block.Nonce = nonce
	block.Hash = blockHash
	fmt.Printf("We've got a mine . Hash:%x nonce:%d\n",block.Hash,block.Nonce)
	_BlockChain.AddBlock(block)
	ShowBlockChainInfo(_BlockChain)
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
	firstBlock := &Block{}
	firstIndex := lastBlock.Index - config.NInterval
	bFirstBlock,err := utils.GetDb().GetBlock(firstIndex)
	err = json.Unmarshal(bFirstBlock,firstBlock)
	if firstIndex < 0 {
		return firstBlock.Diffcult
	}

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

func VerifyBlock(block *Block){
	//verify Hash


}