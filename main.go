package main

import (
	"github.com/kataras/iris"
	"fmt"
	"github.com/EthereumCommonwealth/go-callisto/common"
	"github.com/EthereumCommonwealth/Galileo/models"
)

func main() {
	app := iris.New()

	hash := models.Block{Hash: common.StringToHash("asdasdas"), Miner: common.StringToAddress("asdasdasd")}

	fmt.Println(hash.Hash)
	fmt.Println(hash.Hash.Str())

	app.Run(iris.Addr(":8000"))
}