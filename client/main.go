package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"week1/client/routes"
)

var port = flag.Int("port", 8080, "api client addr")

func main() {
	engine := gin.Default()
	routes.NewRoute(engine)
	err := engine.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("客户端服务开启失败", err)
	}
}
