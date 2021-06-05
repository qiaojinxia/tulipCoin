package config

import "math"

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */


const (
	Version    = "0.0.0.1" //Version of Mine Clinet
	TargetBits = int32(7)
	MaxNone = math.MaxInt64
	DbName = "BlockData"
	BlockHeader = "BH"

	BlockTransactionsID = "TXID"

	BlockTransactions = "TX"
	BlockInfo_Size = "index"
	TransactionPool= "TxPool"
    CHECKSUM_LENGTH = 4

	RewardCoinn = 7
)