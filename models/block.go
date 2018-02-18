package models

import (
	"github.com/jinzhu/gorm"
	"github.com/EthereumCommonwealth/Galileo/common"
	"time"
)

type Block struct {
	gorm.Model
	Hash common.Hash	`gorm:"unique_index:hash_block"`
	ParentHash common.Hash
	Miner string
	TransactionRoot common.Hash
	Difficulty uint
	Number uint
	GasLimit uint
	GasUsed uint
	Timestamp time.Time
	ExtraData string
	Nonce uint
}