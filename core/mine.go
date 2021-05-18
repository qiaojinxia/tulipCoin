package core

import "fmt"

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func Mining(walletAddress []byte){
	lastIndex,preHash := _BlockChain.GetLastBlockIndex()
	block := NewBlock(lastIndex + 1,preHash,walletAddress,"Begin Mine Block,Good Luck!")
	pow := NewProofOfWork(block)
	nonce,blockHash := pow.Run()
	block.Nonce = nonce
	block.Hash = blockHash
	fmt.Printf("We've got a mine . Hash:%x nonce:%d\n",block.Hash,block.Nonce)
	_BlockChain.AddBlock(block)
	ShowBlockChainInfo(_BlockChain)
}