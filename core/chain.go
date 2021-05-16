package core

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */

type BlcokChain []*Block

func(bc *BlcokChain) GetLastBlockChain() *Block{
	return (*bc)[len(*bc)-1]
}

func(bc *BlcokChain) AddBlock(block *Block) {
	*bc = append(*bc, block)
}