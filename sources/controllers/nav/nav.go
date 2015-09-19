package nav

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SidebarTemplateHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "sidebar.tmpl", gin.H{})
}

func HeaderTemplateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "header.tmpl", nil)
}

func FooterTemplateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "footer.tmpl", nil)
}
