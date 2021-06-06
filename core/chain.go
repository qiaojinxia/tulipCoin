package core

import (
	"encoding/json"
	"fmt"
	"log"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */




type BlocksChain struct {
	size int
	Blocks []*Block
}

var _BlockChain *BlocksChain

func init(){
	index,err := utils.GetDb().GetBlockHeight()
	if err != nil{
		panic(err)
	}
	_BlockChain = &BlocksChain{
		size:   index,
		Blocks: make([]*Block,index + 1),
	}
	genesisBlock := CreateGenesisBlock()
	_BlockChain.Blocks[0] = genesisBlock
	res ,err := utils.GetDb().IterAllBlock()
	if err != nil{
		panic(err)
	}
	for _,blockSerialize := range res{
		block := &Block{}
		err := json.Unmarshal(blockSerialize,block)
		if err != nil{
			panic(err)
		}
		_BlockChain.Blocks[block.Index] = block
	}
	fmt.Println(res)
}

func(bc *BlocksChain) GetLastBlockChain() *Block{
	return bc.Blocks[bc.size]
}

func(bc *BlocksChain) FindTransactionByTxID(txID []byte) *Transaction{
	b_tx,err := utils.GetDb().GetTransactionByTxID(txID)
	if err != nil{
		panic(err)
	}
	tx := &Transaction{}
	err = json.Unmarshal(b_tx,tx)
	if err != nil{
		panic(err)
	}
	return tx
}

func(bc *BlocksChain) AddBlock(block *Block){
	if !ValidBlock(bc.GetLastBlockChain(),block){
		return
	}
	fmt.Printf("New Block To add BlockChain Success!%v\n",*block)
	bc.Blocks = append(bc.Blocks, block)
	bc.size ++
	blockSerialize,err := json.Marshal(block)
	if err != nil{
		log.Panic(err)
	}
	err = utils.GetDb().StoreBlock(block.Index,blockSerialize)
	if err != nil{
		log.Panic(err)
	}
	transactionsID := make([][]byte,0,len(block.Transactions))
	for _,tx := range block.Transactions{
		txSerlalize,err := json.Marshal(tx)
		transactionsID = append(transactionsID, tx.TxID)
		if err != nil{
			log.Panic(err)
		}
		err = utils.GetDb().StoreTransaction(tx.TxID,txSerlalize)
		if err != nil{
			log.Panic(err)
		}
	}
	tranText,err := json.Marshal(transactionsID)
	if err != nil{
		panic(utils.ConverErrorWarp(err,""))
	}
	err = utils.GetDb().StoreTransactionsID(block.MRoot,tranText)
	if err != nil{
		panic(utils.ConverErrorWarp(err,""))
	}
	err = utils.GetDb().StoreBlockHeight(utils.ToBytes(block.Index))
	if err != nil{
		log.Panic(err)
	}
}

func(bc *BlocksChain) GetLastBlockIndex() (int,[]byte){
	return bc.size,bc.Blocks[bc.size].Hash
}

func ShowBlockChainInfo(bc *BlocksChain){
	fmt.Printf("※※※※※※※ Print All BlockChain Info ※※※※※※※\n")
	fmt.Printf("※※※※※※※※ All Block  Num %d  ※※※※※※※\n",bc.size + 1)

	for _,block := range bc.Blocks{
		if block.Index == 0{
			continue
		}
		fmt.Printf("Block  Index : %d\n",block.Index)
		fmt.Printf("Block  Hash : %x\n",block.Hash)
		fmt.Printf("Block  Nonce : %d\n",block.Nonce)
		fmt.Printf("Prev Block Hash : %x\n",block.PreviousHash)
		var txsID [][]byte
		bTxID,err := utils.GetDb().GetTransactionsID(block.MRoot)
		if err != nil{
			panic(utils.MarshalErrorWarp(err,""))
		}
		err = json.Unmarshal(bTxID,&txsID)
		if err != nil{
			panic(utils.MarshalErrorWarp(err,""))
		}
		txs := make([]Transaction,0)
		for _,txID := range txsID{
			tx := _BlockChain.FindTransactionByTxID(txID)
			txs = append(txs, *tx)
		}

		fmt.Printf("Block Transactions : %v\n",txs)

		fmt.Printf("Block TimeStamp : %d\n",block.TimeStamp)

		fmt.Printf("==================\n")
	}

}
func ValidBlock(preBlock,nowBlock *Block) bool{
	nowBlock.PreviousHash = preBlock.Hash
	pow := NewProofOfWork(nowBlock)
	return pow.Validate()
}

