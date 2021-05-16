package core

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */
import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"main/config"
	"main/utils"
	"strconv"
)
//BlotDB的管理类
type BlockChainDB struct {
	db *bolt.DB //db
}

func NewBlockChainDb() (*BlockChainDB,error){
	mdb, err := bolt.Open("block" , 0644, nil)
	if err != nil {
		return nil,err
	}
	err = mdb.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte(config.DbName))
		tx.CreateBucket([]byte(config.BlockInfo))
		return nil
	})
	return &BlockChainDB{
		db:mdb,
	},nil
}


func(m *BlockChainDB) PutBlockSize(val []byte) (err error){
	err  = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BlockInfo))
		return b.Put([]byte(config.BlockInfo_Size), val)
	})
	return
}

func(m *BlockChainDB) GetBlockSize() (index int,err error){
	err  = m.db.Update(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte(config.BlockInfo))
	tmp := b.Get([]byte(config.BlockInfo_Size))
	index = int(utils.BytesToInt64(tmp))
	return err
	})
	return
}

func (m *BlockChainDB) Add(bucketName string, val []byte) (id uint64, err error) {
	err = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id, _ = b.NextSequence() //sequence uint64
		bBuf := fmt.Sprintf("%d", id)
		return b.Put([]byte(bBuf), val)
	})
	return
}


func (m *BlockChainDB) Select(bucketName string) (res map[int64][]byte,err error) {
	blockJsonMap := make(map[int64][]byte)
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			log.Printf("key=%s, value=%s\n", string(k), v)
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