package main

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"github.com/EthereumCommonwealth/Galileo/common"
	"github.com/EthereumCommonwealth/Galileo/models"
	"github.com/labstack/echo"
	"github.com/EthereumCommonwealth/Galileo/api"
	"github.com/labstack/echo/middleware"
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
	e.Debug = true

	// Server Middlewares

	e.Pre(middleware.AddTrailingSlash())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/addr/:address/", api.GetAddressDetails)

	e.Logger.Fatal(e.Start(":8000"))
}