package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/EthereumCommonwealth/Galileo/models"
)

func GetAddressDetails(c echo.Context) error {
	addr := models.Address{}
	return c.JSON(http.StatusOK, addr)
}
