package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func App(r *gin.Engine) {
	r.GET("/:text", func(c *gin.Context) {
		text := c.Param("text")
		data := GetDataset()
		result := Convert(data, text)
		c.String(http.StatusOK, result)
	})
}
