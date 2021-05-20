package utils

import (
	"log"
	"main/config"
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
