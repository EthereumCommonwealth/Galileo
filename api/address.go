package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/EthereumCommonwealth/Galileo/models"
	"github.com/EthereumCommonwealth/go-callisto/common"
	"fmt"
)

var messages = map[string]string{
	"invalid_address": "The address: %s is not a valid Ethereum-like address",
}

func GetAddressDetails(c echo.Context) error {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": fmt.Sprintf(messages["invalid_address"], address),
		})
	}
	addr := models.Address{}
	return c.JSON(http.StatusOK, addr)
}
