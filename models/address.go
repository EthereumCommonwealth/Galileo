package models

import (
	"github.com/jinzhu/gorm"
	"github.com/EthereumCommonwealth/go-callisto/common"
)

type Address struct {
	gorm.Model
	Address common.Address
	Type    string
}