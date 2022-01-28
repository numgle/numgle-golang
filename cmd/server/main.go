package main

import (
	"github.com/gin-gonic/gin"
	"github.com/numgle/numgle-golang/internal"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	internal.App(r)
	r.Run(":3000")
}
