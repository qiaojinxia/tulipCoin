package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2021/5/27.
 *  email:<115882934@qq.com>
 */

func LoadPrivateKey(){
	os.Open("")
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}


func LoadWallet(path string) *BitcoinKeys{
	file,err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err!= nil{
		panic(err)
	}
	content := make([]byte,0)
	defer file.Close()
	buff := make([]byte,1024)
	for {
		n,err := file.Read(buff)
		if err != nil{
			if err == io.EOF{
				break
			}
			panic(err)
		}
		content = append(content, buff[:n]...)
	}
	key := strings.Split(string(content),";")
	wallet := &BitcoinKeys{}
	wallet.Init(key[0],key[1])
	return wallet
}