package core

/**
 * Created by @CaomaoBoy on 2021/6/1.
 *  email:<115882934@qq.com>
 */
import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"sort"
)
type MerkleTree struct {
	Root       *MerkleNode
	merkleRoot []byte
	leaves     []*MerkleNode
}

type MerkleNode struct {
	left    *MerkleNode
	right   *MerkleNode
	prev    *MerkleNode
	isleaf  bool
	content []byte
}

type Proof struct {
	merklePath [][]byte
	index int
}

//Get the adjacent hash  from node(txID) to the merklet root
func(mt *MerkleTree) GenerateMerkleProof(txID []byte) *Proof{
	index := binarySearch(mt.leaves,txID)
	if index == -1{
		return nil
	}
	levelCount := int(math.Ceil(math.Log2(float64(len(mt.leaves)))))
	merklePath := make([][]byte,0,levelCount)
	var cur = mt.leaves[index]
	//get uncle node
	for cur.prev != nil{
		if bytes.Compare(cur.prev.left.content,cur.content) == 0{
			if cur.left == nil && cur.right == nil {
				merklePath = append(merklePath, []byte{})
			}
			merklePath = append(merklePath, cur.prev.right.content)
		}else if bytes.Compare(cur.prev.right.content,cur.content) == 0{
			merklePath = append(merklePath, cur.prev.left.content)
			if cur.left == nil && cur.right == nil {
				merklePath = append(merklePath, []byte{})
			}
		}
		cur = cur.prev
	}
	proof := &Proof{
		index: index,
		merklePath:merklePath,
	}
	return proof
}

func(mt *MerkleTree) VerifyMerklet(merkleRoot []byte,dataHash []byte,proof *Proof) bool {
	var hash []byte
	m := proof.index % 2
	proof.merklePath[m] = dataHash
	hash = append(hash, proof.merklePath[0]...)
	for i:=1;i<len(proof.merklePath);i+=1{
		if i % 2 == 0{
			hash = append(proof.merklePath[i],hash...)
		}else{
			hash = append(hash, proof.merklePath[i]...)
		}
		hashData := sha256.Sum256(hash)
		hash1 :=  sha256.Sum256(hashData[:])
		hash = hash1[:]
	}
	if bytes.Compare(merkleRoot,hash) == 0{
		return true
	}
	return false
}

func min(x,y int ) int{
	if x < y{
		return x
	}
	return y
}


func binarySearch(content []*MerkleNode,txID []byte) int{
	left := 0
	right := len(content)
	b1 := big.Int{}
	b1.SetBytes(txID)
	for left <= right{
		mid := left + (right - left) / 2
		b2:= big.Int{}
		b2.SetBytes(content[mid].content)
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
	for _,v := range hashData{
		node := NewMerkleNode(nil,nil,v)
		nodes = append(nodes,node)
	}
	content := nodes
	sort.SliceStable(nodes, func(i, j int) bool {
		b1 := big.Int{}
		b2 := big.Int{}
		b1.SetBytes(nodes[i].content)
		b2.SetBytes(nodes[j].content)
		return b1.Cmp(&b2) == -1
	})
	for len(nodes)>1{
		var tmpNode []*MerkleNode
		for i:=0;i<len(nodes);i+=2{
			i2 := min(i+1,len(nodes) -1)
			node := NewMerkleNode(nodes[i],nodes[i2],nil)
			tmpNode = append(tmpNode, node)
		}
		nodes = tmpNode
	}
	return &MerkleTree{Root: nodes[0], leaves:content,merkleRoot: nodes[0].content}
}

func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode{
	var mnode = new(MerkleNode)
	if left == nil && right == nil{
		mnode = &MerkleNode{
			left:    nil,
			right:   nil,
			isleaf:  true,
			content: data,
		}
	}else{
		childsData := append(left.content,right.content...)
		hashData := sha256.Sum256(childsData)
		hash :=  sha256.Sum256(hashData[:])
		mnode = &MerkleNode{
			left:    left,
			right:   right,
			prev:    nil,
			content: hash[:],
		}
		left.prev = mnode
		right.prev = mnode
	}
	return mnode
}