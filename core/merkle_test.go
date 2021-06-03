package core

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_MerkleTree(t *testing.T) {
	c1 := sha256.Sum256([]byte("1231231SADAdsfSD2321"))
	c2 := sha256.Sum256([]byte("123123SADA123SD12321"))
	c3 := sha256.Sum256([]byte("12312ASDSAD213SAD312321"))
	c4 := sha256.Sum256([]byte("12312AasdasdSDSADSAD312321"))
	tree := NewMerkleTree([][]byte{
		c1[:],
		c2[:],
		c3[:],
		c4[:],
	})
	for _,v := range tree.Verify(c3[:]){
		fmt.Printf("%x\n",v)
	}
	fmt.Printf("Merkle Root %x\n",tree.merkleRoot)

	//fmt.Printf("%x\n",tree.Root.left.left.hashData)
	//fmt.Printf("%x\n",tree.Root.left.right.hashData)
	//fmt.Printf("%x\n",tree.Root.hashData)
}