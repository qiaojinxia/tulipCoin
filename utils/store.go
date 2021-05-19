package utils

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */
import (
	"fmt"
	"github.com/boltdb/bolt"
	"main/config"
	"strconv"
)
//BlotDB
type BlockChainDB struct {
	db *bolt.DB //db
}

func NewBlockChainDb() (*BlockChainDB,error){
	mdb, err := bolt.Open("block.db" , 0644, nil)
	if err != nil {
		return nil,err
	}
	err = mdb.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte(config.DbName))
		tx.CreateBucket([]byte(config.BlockHeader))
		tx.CreateBucket([]byte(config.BlockTransactions))
		return nil
	})
	return &BlockChainDB{
		db:mdb,
	},nil
}


func(m *BlockChainDB) StoreBlockHeight(val []byte) (err error){
	err  = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BlockHeader))
		return b.Put([]byte(config.BlockInfo_Size), val)
	})
	return
}

func(m *BlockChainDB) GetBlockSize() (index int,err error){
	err  = m.db.View(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte(config.BlockHeader))
	tmp := b.Get([]byte(config.BlockInfo_Size))
	index = int(BytesToInt64(tmp))
	return err
	})
	return
}

func (m *BlockChainDB) StoreBlock(blockIndex int64, data []byte) (err error) {
	err = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.DbName))
		bIndex := fmt.Sprintf("%d", blockIndex)
		return b.Put([]byte(bIndex), data)
	})
	return
}

func(m *BlockChainDB) StoreTransaction(TxID []byte,data []byte) (err error) {
	err = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BlockTransactions))
		return b.Put(TxID, data)
	})
	return
}

func(m *BlockChainDB) GetTransactionByTxID(TxID []byte) (transaction []byte,err error){
	err = m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BlockTransactions))
		transaction = b.Get(TxID)
		return err
	})
	return
}

func(m *BlockChainDB) GetAllTransactions() (transactions [][]byte,err error){
	transactions = make([][]byte,0)
	err = m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BlockTransactions))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			transactions = append(transactions, v)
			return nil
		})
		return nil
	})
	return
}


func (m *BlockChainDB) IterAllBlock() (res map[int64][]byte,err error) {
	blockJsonMap := make(map[int64][]byte)
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.DbName))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			blockIndex,err := strconv.ParseInt(string(k),10,64)
			if err != nil{
				panic(err)
			}
			blockJsonMap[blockIndex] = v
			return nil
		})
		return nil
	})
	return blockJsonMap,nil
}


func (m *BlockChainDB) Close() error {
	return m.db.Close()
}