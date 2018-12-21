package main

import (
	"log"
	"wechat/pkg/bootstrap"
	"wechat/pkg/logs"
)

const port = ":8080"

func main()  {
	logs.ZapLogInitialize()
	framework := bootstrap.NewConsole().FrameworkInit().RouterInit()
	err := framework.Run(port)
	if err != nil {
		log.Fatal(err.Error())
	}
}