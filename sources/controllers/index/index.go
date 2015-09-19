package index

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.File("static/index.html")
}
