package core

import "sync"

/**
 * Created by @CaomaoBoy on 2021/5/25.
 *  email:<115882934@qq.com>
 */
var once sync.Once
var CtxPool *CtxMemPool
func GetCtxPool() *CtxMemPool {
	once.Do(
		func() {
			CtxPool = &CtxMemPool{Transactions: make([]*CTxMemPoolEntry,0)}
		},
	)
	return CtxPool

}

type CTxMemPoolEntry struct {
	CTransactionRef *Transaction
	nFee float64 //Fee
	nUsageSize float64 //Transaction Mem Use
	nTime int64 //Time when the transaction joins the blockChain
	entryHeight int // Height of transaction joining  blockChain
	spendsCoinbase bool //prev Transaction is Coinbase
}

type CtxMemPool struct {
	Transactions []*CTxMemPoolEntry
}

func(cmp *CtxMemPool) AddTxToPool(ctxMemEntry *CTxMemPoolEntry){
	cmp.Transactions = append(cmp.Transactions, ctxMemEntry)
}

func(cmp *CtxMemPool) PopTx() *CTxMemPoolEntry{
	tmp := cmp.Transactions[len(cmp.Transactions)-1]
	cmp.Transactions = cmp.Transactions[:len(cmp.Transactions)-1]
	return tmp
}