package core

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_MerkleTree(t *testing.T) {
	c1 := sha256.Sum256([]byte("1231231SADASD2321"))
	c2 := sha256.Sum256([]byte("123123SADASD12321"))
	c3 := sha256.Sum256([]byte("12312ASDSADSAD312321"))
	tree := NewMerkleTree([][]byte{
		c1[:],
		c2[:],
		c3[:],
		c3[:],
		c3[:],
	})
	fmt.Printf("%x\n",tree.Root.left.left.hashData)
	fmt.Printf("%x\n",tree.Root.left.right.hashData)
	fmt.Printf("%x\n",tree.Root.hashData)
}