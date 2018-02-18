package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/EthereumCommonwealth/go-callisto/common"
)

type Block struct {
	gorm.Model
	Hash            common.Hash `gorm:"unique_index:hash_block"`
	ParentHash      common.Hash
	Miner           common.Address
	TransactionRoot common.Hash
	Difficulty      uint
	Number          uint
	GasLimit        uint
	GasUsed         uint
	Timestamp       time.Time
	ExtraData       string
	Nonce           uint
}
