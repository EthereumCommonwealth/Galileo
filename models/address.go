package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/EthereumCommonwealth/go-callisto/common"
)

type AddressType uint8

const (
	NormalAddress   AddressType = 1
	ContractAddress AddressType = 2
)

type Address struct {
	gorm.Model
	Address string      `gorm:"type:varchar(50);unique_index:address"`
	Type    AddressType
}

func NewAddress(address string) *Address {
	return &Address{Address: address}
}

func (a *Address) Save(db *gorm.DB) {
	if !db.NewRecord(a) {
		return
	}

	db.Create(a)
}

func (a *Address) GetAddressDetails(db *gorm.DB) {
	db.Where(&Address{Address: a.Address}).First(a)
}

func (a *Address) IsValid() bool {
	return common.IsHexAddress(a.Address)
}

func (a *Address) AddressType() string {
	switch a.Type {
	case NormalAddress:
		return "normal"
	case ContractAddress:
		return "contract"
	default:
		return "normal"
	}
}

func (a *Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		gorm.Model
		Address string
		Type    string
	}{
		a.Model,
		a.Address,
		a.AddressType(),
	})
}
