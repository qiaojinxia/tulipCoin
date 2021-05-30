package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"log"
	"main/config"
	"math/big"
)

/**
 * Created by @CaomaoBoy on 2021/5/19.
 *  email:<115882934@qq.com>
 */
func WalletAddressToPublicKeyHash(walletAddress []byte) []byte{
	fullHash := Base58Decode(walletAddress)
	if len(fullHash) != 25 {
		log.Panic("not valid Wallet!")
	}
	return fullHash[1:len(fullHash)- config.CHECKSUM_LENGTH]

}

func GenerateLockScript(walletAddress []byte) string{
	return fmt.Sprintf("OP_DUP OP_HASH160 %x OP_EQUALVERIFY OP_CHECKSIG ",WalletAddressToPublicKeyHash(walletAddress))
}


func GenerateLockScriptByPublicHash(publicKey []byte) string{
	return fmt.Sprintf("OP_DUP OP_HASH160 %x OP_EQUALVERIFY OP_CHECKSIG ",GeneratePublicKeyHash(publicKey))
}

func Verify(sig,publicKey[]byte,dataHash []byte) bool{
	r := big.Int{}
	s := big.Int{}

	r.SetBytes(sig[:len(sig)/2])
	s.SetBytes(sig[len(sig)/2:])

	X := big.Int{}
	Y := big.Int{}

	X.SetBytes(publicKey[:len(publicKey)/2])
	Y.SetBytes(publicKey[len(publicKey)/2:])

	pubKeyOrigin := ecdsa.PublicKey{elliptic.P256(), &X, &Y}
	return ecdsa.Verify(&pubKeyOrigin,dataHash, &r, &s)
}

