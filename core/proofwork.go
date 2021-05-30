package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"main/config"
	"main/utils"
	"math/big"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */


type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork{
	targrt := big.NewInt(1)
	targrt.Lsh(targrt,uint(256-config.TargetBits))
	pow := &ProofOfWork{b,targrt}
	return pow
}

func(pow *ProofOfWork) prepareData(nonce int64) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PreviousHash,
			pow.block.MRoot,
			utils.ToBytes(pow.block.Index),
			utils.ToBytes(pow.block.TimeStamp),
			[]byte(pow.block.Version),
			utils.ToBytes(config.TargetBits),
			utils.ToBytes(nonce),
		},
		[]byte{},
	)
	return data
}

func(pow *ProofOfWork) Run() (int64,[]byte){
	var hashInt big.Int
	var hash [32]byte
	var nonce int64
	fmt.Printf("Mine the block containing \n")
	for nonce < config.MaxNone {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		fmt.Printf("\n")
		if hashInt.Cmp(pow.target) == -1{
			break
		}else{
			nonce ++
		}
	}
	return nonce,hash[:]
}

func(pow *ProofOfWork) Validate() bool{
	//Validate Diffculty
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}