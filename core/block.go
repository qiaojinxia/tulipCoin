package core

import "main/config"

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */

type Block struct {
	Index int `json:"index"`
	PreviousHash string `json:"previous_hash"`
	Memorials  []*Memorial `json:"memorial"`
	TimeStamp int64 `json:"time_stamp"`
	Data string `json:"data"`
	Hash int64 `json:"hash"`
	Version string `json:"version"`
}

func NewBlock() *Block{
	return &Block{
		Index:        0,
		PreviousHash: "",
		TimeStamp:    0,
		Data:         "",
		Hash:         0,
		Version:      config.Version,
	}
}

//创造创世区块
func CreateGenesisBlock() *Block{
	return &Block{
		Index:        1,
		PreviousHash: "",
		TimeStamp:    0,
		Data:         "it's my time to create word!",
		Hash:         0,
		Version:      config.Version,
	}
}