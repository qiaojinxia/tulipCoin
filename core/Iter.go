package core

import (
	"encoding/json"
	"log"
)

/**
 * Created by @CaomaoBoy on 2021/5/19.
 *  email:<115882934@qq.com>
 */

type IIterator interface {
	HasNext() bool
	Next() *Transaction
}

type TransactionIter struct {
	index int
	transactions []*Transaction
}

func (t *TransactionIter) HasNext() bool {
	if t.index < len(t.transactions) {
		return true
	}
	return false
}

func (t *TransactionIter) Next() *Transaction {
	if t.HasNext() {
		n := t.index
		t.index ++
		return t.transactions[n]
	}
	return nil
}

func NewTransactionIter(bTransacions [][]byte) IIterator {
	txs := make([]*Transaction,0,len(bTransacions))
	for _,tdata := range bTransacions{
		tx := &Transaction{}
		err := json.Unmarshal(tdata,tx)
		if err != nil{
			log.Panic(err)
		}
		txs = append(txs, tx)
	}
	return &TransactionIter{
		index:        0,
		transactions: txs,
	}
}

