package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	localProxy    string
	clientDistDir string
	port          string
)

func init() {
	flag.StringVar(&localProxy, "proxy", "", "proxy address, eg: http://localhost:1080")
	flag.StringVar(&clientDistDir, "client-dir", "dist/client", "dist/client directory path")
	flag.StringVar(&port, "port", ":9999", "server listen port")
}

func main() {
	flag.Parse()

	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {

	})
}
