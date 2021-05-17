package core

import (
	"encoding/json"
	"fmt"
	"main/config"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */




type BlocksChain struct {
	size int
	Blocks []*Block
	*BlockChainDB
}

var _BlockChain *BlocksChain

func init(){
	blockChainDB,err := NewBlockChainDb()
	if err != nil{
		panic(err)
	}
	index,err := blockChainDB.GetBlockSize()
	if err != nil{
		panic(err)
	}
	_BlockChain = &BlocksChain{
		size:   index,
		Blocks: make([]*Block,index + 1),
	}
	_BlockChain.BlockChainDB = blockChainDB
	genesisBlock := CreateGenesisBlock()
	_BlockChain.Blocks[0] = genesisBlock
	res ,err := blockChainDB.Select(config.DbName)
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

func(bc *BlocksChain) AddBlock(block *Block){
	if !ValidBlock(bc.GetLastBlockChain(),block){
		return
	}
	fmt.Printf("New Block To add BlockChain Success!%v\n",*block)
	bc.Blocks = append(bc.Blocks, block)
	bc.size ++
	blockSerialize,err := json.Marshal(block)
	if err != nil{
		panic(err)
	}
	bc.Add(config.DbName,blockSerialize)
	bc.PutBlockSize(utils.ToBytes(block.Index))
}

func(bc *BlocksChain) GetLastBlockIndex() (int,[]byte){
	return bc.size,bc.Blocks[bc.size].Hash
}

func ShowBlockChainInfo(bc *BlocksChain){
	fmt.Printf("※※※※※※※ Print All BlockChain Info ※※※※※※※\n")
	fmt.Printf("※※※※※※※※ All Block  Num %d  ※※※※※※※\n",bc.size + 1)

	for _,block := range bc.Blocks{
		fmt.Printf("Block  Index : %d\n",block.Index)
		fmt.Printf("Block  Hash : %x\n",block.Hash)
		fmt.Printf("Block  Nonce : %d\n",block.Nonce)
		fmt.Printf("Prev Block Hash : %x\n",block.PreviousHash)
		fmt.Printf("Block TimeStamp : %d\n",block.TimeStamp)
		fmt.Printf("==================\n")
	}

}
func ValidBlock(preBlock,nowBlock *Block) bool{
	nowBlock.PreviousHash = preBlock.Hash
	pow := NewProofOfWork(nowBlock)
	return pow.Validate()
}

