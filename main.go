package main

import "go-blog/router"

func main()  {
	server := router.RouterInit()
	server.Run(":8888")

}
