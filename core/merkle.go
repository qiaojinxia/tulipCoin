package core

/**
 * Created by @CaomaoBoy on 2021/6/1.
 *  email:<115882934@qq.com>
 */
import (
	"crypto/sha256"
	"math/big"
	"sort"
)

type MerkleTree struct {
	Root *MerkleNode
	merkleRoot []byte
	content []*MerkleNode
}

type MerkleNode struct {
	left *MerkleNode
	right *MerkleNode
	prev *MerkleNode
	isleft bool
	hashData []byte
}

func(mt *MerkleTree) Verify(txID []byte) [][]byte{
	index := binarySearch(mt.content,txID)
	if index == -1{
		return nil
	}
	merklePath := make([][]byte,0,index)
	var cur = mt.content[index].prev
	for cur != nil{
		merklePath = append(merklePath, cur.left.hashData)
		merklePath = append(merklePath, cur.right.hashData)
		cur = cur.prev
	}
	return merklePath
}

func binarySearch(content []*MerkleNode,txID []byte) int{
	left := 0
	right := len(content)
	b1 := big.Int{}
	b1.SetBytes(txID)
	for left <= right{
		mid := left + (right - left) / 2
		b2:= big.Int{}
		b2.SetBytes(content[mid].hashData)
		bigOrSmall := b1.Cmp(&b2)
		if bigOrSmall == 1{	//b1 > b2
			left = mid + 1
		}else if bigOrSmall == 0{ //b1 = b2
			return mid
		}else{//b1< b2
			right = mid -1
		}
	}
	return -1
}

func NewMerkleTree(hashData [][]byte) *MerkleTree{
	var nodes []*MerkleNode
	if len(hashData) % 2 != 0{
		hashData = append(hashData, hashData[len(hashData)-1])
	}
	for _,v := range hashData{
		node := NewMerkleNode(nil,nil,v)
		nodes = append(nodes,node)
	}
	content := nodes
	sort.SliceStable(nodes, func(i, j int) bool {
		b1 := big.Int{}
		b2 := big.Int{}
		b1.SetBytes(nodes[i].hashData)
		b2.SetBytes(nodes[j].hashData)
		return b1.Cmp(&b2) == -1
	})
	for ;len(nodes)>1;{
		var tmpNode []*MerkleNode
		for i:=0;i<len(nodes);i+=2{
			node := NewMerkleNode(nodes[i],nodes[i+1],nil)
			nodes[i].prev = node
			nodes[i+1].prev = node
			tmpNode = append(tmpNode, node)
		}
		if len(tmpNode) != 1 && len(tmpNode) % 2 != 0{
			tmpNode = append(tmpNode, tmpNode[len(tmpNode)-1])
		}
		nodes = tmpNode
	}
	return &MerkleTree{Root: nodes[0],content:content,merkleRoot: nodes[0].hashData}
}

func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode{
	var mnode = new(MerkleNode)
	if left == nil && right == nil{
		mnode = &MerkleNode{
			left:     nil,
			right:    nil,
			isleft :true,
			hashData: data,
		}
	}else{
		childsData := append(left.hashData,right.hashData...)
		hashData := sha256.Sum256(childsData)
		hash :=  sha256.Sum256(hashData[:])
		mnode = &MerkleNode{
			left:     left,
			right:    right,
			prev:     nil,
			hashData: hash[:],
		}
	}
	return mnode
}

