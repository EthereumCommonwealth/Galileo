package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/EthereumCommonwealth/Galileo/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

var messages = map[string]string{
	"invalid_address": "The address: %s is not a valid Ethereum-like address",
}

func GetAddressDetails(c echo.Context) error {
	address := models.NewAddress(c.Param("address"))
	if !address.IsValid() {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": fmt.Sprintf(messages["invalid_address"], address.Address),
		})
	}
	address.GetAddressDetails(c.Get("db").(*gorm.DB))
	return c.JSON(http.StatusOK, address)
}
