package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
	"log"
	"main/config"
)

/**
 * Created by @CaomaoBoy on 2021/5/17.
 *  email:<115882934@qq.com>
 */

const VERSION = byte(0x00)



type BitcoinKeys struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func GetBitcoinKeys() *BitcoinKeys {
	b := &BitcoinKeys{nil, nil}
	b.newKeyPair()
	return b
}

//Elliptic curve generate  PublicKey
func (b *BitcoinKeys) newKeyPair(){
	curve := elliptic.P256()
	var err error
	b.PrivateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil{
		log.Panic(err)
	}
	b.PublicKey = append(b.PrivateKey.PublicKey.X.Bytes(), b.PrivateKey.PublicKey.Y.Bytes()...)
}

func GeneratePublicKeyHash(publicKey []byte) []byte {
	sha256PubKey := sha256.Sum256(publicKey)
	r := ripemd160.New()
	r.Write(sha256PubKey[:])
	ripPubKey := r.Sum(nil)
	return ripPubKey
}

//获取地址
func (b *BitcoinKeys) GetAddress() []byte {
	//1.ripemd160(sha256(publickey))
	ripPubKey := GeneratePublicKeyHash(b.PublicKey)
	//2.add one byte version info to head versionPublickeyHash
	versionPublickeyHash := append([]byte{VERSION}, ripPubKey[:]...)
	//3.checksumHash = sha256(sha256(versionPublickeyHash))[:4]
	tailHash := CheckSumHash(versionPublickeyHash)
	//4.join bytes slic  versionPublickeyHash + checksumHash
	finalHash := append(versionPublickeyHash, tailHash...)
	//5.Do base58 encryption
	address := Base58Encode(finalHash)
	return address
}
func (b *BitcoinKeys) GetPrivateKey() []byte{
	ecder,err := x509.MarshalECPrivateKey(b.PrivateKey)
	if err != nil{
		log.Panic(err)
	}
	return Base58Encode(ecder)
}
func CheckSumHash(versionPublickeyHash []byte) []byte {
	versionPublickeyHashSha1 := sha256.Sum256(versionPublickeyHash)
	versionPublickeyHashSha2 := sha256.Sum256(versionPublickeyHashSha1[:])
	tailHash := versionPublickeyHashSha2[:config.CHECKSUM_LENGTH]
	return tailHash
}

func IsVaildBitcoinAddress(address string) bool {
	adddressByte := []byte(address)
	fullHash := Base58Decode(adddressByte)
	if len(fullHash) != 25 {
		return false
	}
	prefixHash := fullHash[:len(fullHash)-config.CHECKSUM_LENGTH]
	tailHash := fullHash[len(fullHash)-config.CHECKSUM_LENGTH:]
	tailHash2 := CheckSumHash(prefixHash)
	if bytes.Compare(tailHash, tailHash2[:]) == 0 {
		return true
	} else {
		return false
	}
}

func(b *BitcoinKeys) Sign(data []byte) (string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, b.PrivateKey, data)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var bf bytes.Buffer
	w := gzip.NewWriter(&bf)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(bf.Bytes()), nil
}

