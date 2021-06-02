package core

import (
	"crypto/sha256"
	"math/big"
	"sort"
)

type MerkleTree struct {
	Root *MerkleNode
	merkleRoot []byte
}

type MerkleNode struct {
	left *MerkleNode
	right *MerkleNode
	prev *MerkleNode
	hashData []byte
}

func NewMerkleTree(hashData [][]byte) *MerkleTree{
	var nodes []MerkleNode
	if len(hashData) % 2 != 0{
		hashData = append(hashData, hashData[len(hashData)-1])
	}
	for _,v := range hashData{
		node := NewMerkleNode(nil,nil,v)
		nodes = append(nodes,*node)
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		b1 := big.Int{}
		b2 := big.Int{}
		b1.SetBytes(nodes[i].hashData)
		b2.SetBytes(nodes[j].hashData)
		return b1.Cmp(&b2) == -1
	})
	for ;len(nodes)>1;{
		var tmpNode []MerkleNode
		for i:=0;i<len(nodes);i+=2{
			node := NewMerkleNode(&nodes[i],&nodes[i+1],nil)
			nodes[i].prev = node
			nodes[i+1].prev = node
			tmpNode = append(tmpNode, *node)
		}
		if len(tmpNode) != 1 && len(tmpNode) % 2 != 0{
			tmpNode = append(tmpNode, tmpNode[len(tmpNode)-1])
		}
		sort.SliceStable(tmpNode, func(i, j int) bool {
			b1 := big.Int{}
			b2 := big.Int{}
			b1.SetBytes(nodes[i].hashData)
			b2.SetBytes(nodes[j].hashData)
			return b1.Cmp(&b2) == -1
		})
		nodes = tmpNode
	}
	return &MerkleTree{Root: &nodes[0],merkleRoot: nodes[0].hashData}
}

func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode{
	var mnode = new(MerkleNode)
	if left == nil && right == nil{
		mnode = &MerkleNode{
			left:     nil,
			right:    nil,
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

