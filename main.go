package main

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"github.com/EthereumCommonwealth/Galileo/common"
	"github.com/EthereumCommonwealth/Galileo/models"
	"github.com/labstack/echo"
)

func main() {
	// Galileo Settings

	var galileoSetting common.GalileoSetting
	err := envconfig.Process("galileo", &galileoSetting)
	if err != nil {
		log.Fatal(err)
	}

	// Database

	db := models.GetDBConnection(galileoSetting)
	models.Migrate(db)

	// Server

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}