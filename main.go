package main

import (
	"go-blog/bootstrap"
	"go-blog/bootInject"
)

func main()  {
	boot := bootstrap.Init(
		bootInject.BootDatabase(),
		bootInject.BootGin(),
	)
	boot.Router.Run(":8888")
}
