package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Run(iris.Addr(":8000"))
}