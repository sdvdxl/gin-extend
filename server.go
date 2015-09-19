// main.go
package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sdvdxl/gin-extend/sources/controllers/index"
	"github.com/sdvdxl/gin-extend/sources/controllers/nav"
	"github.com/sdvdxl/gin-extend/sources/util"
	"github.com/sdvdxl/gin-extend/sources/util/db"
	"github.com/sdvdxl/gin-extend/sources/util/log"
	"github.com/sdvdxl/gin-extend/sources/util/render"
	"github.com/sdvdxl/gin-extend/sources/util/session"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	g *gin.Engine
)

func init() {
	//初始化gin处理引擎
	//	gin.SetMode(gin.ReleaseMode)
	g = gin.New()
	g.Use(HandlerError())

	{
		funcMap := template.FuncMap{"Equals": func(v1, v2 interface{}) bool {
			log.Logger.Debug("invoke function Equals")
			return v1 == v2
		}}
		tmp := template.New("myTemplate")
		templatePages := TemplatesFinder("templates")
		tmp.Funcs(funcMap).ParseFiles(templatePages...)

		g.SetHTMLTemplate(tmp)

		{ //这三个顺序不能变更,否则得不到正常处理
			//先设置/读取session信息
			g.Use(sessions.Sessions("my_session", session.SessionStore))

			//然后其他

			//最后处理静态文件
			g.Use(static.ServeRoot("/", "static")) // static files have higher priority over dynamic routes

		}

	}
}

func HandlerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				trace := make([]byte, 10240)
				runtime.Stack(trace, true)
				log.Logger.Error("%s, \n%s", err, trace)

				if strings.HasSuffix(c.Request.URL.Path, ".html") {
					c.HTML(http.StatusInternalServerError, "500.tmpl", nil)
				} else {
					r := render.New(c)
					r.JSON(util.JsonResult{Msg: "系统错误"})
				}
			}
		}()
		c.Next()
	}
}

func main() {
	defer log.Logger.Close()
	defer db.Close()

	log.Logger.Info("starting server...")

	//==================================   nav top sidebar  ======================================
	{
		g.Any("/sidebar.html", nav.SidebarTemplateHandler)
		g.Any("/header.html", nav.HeaderTemplateHandler)
		g.Any("/footer.html", nav.FooterTemplateHandler)
	}

	//==================================   404, index  ======================================
	{
		//404
		g.NoRoute(func(c *gin.Context) {
			log.Logger.Debug("page [%v] not found, redirect to 404.html", c.Request.URL.Path)
			if strings.HasSuffix(c.Request.URL.Path, ".html") {
				c.HTML(http.StatusNotFound, "404.tmpl", nil)
			} else {
				r := render.New(c)
				r.JSON(util.JsonResult{Msg: "请求资源不存在"})
			}
		})

		//首页
		g.Any("/", index.IndexHandler)
	}

	log.Logger.Info("server started ")
	g.Run(":8085")
}

func TemplatesFinder(templateDirName string) []string {
	templatePages := make([]string, 0, 10)
	filepath.Walk(templateDirName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			templatePages = append(templatePages, path)
		}
		return nil
	})

	log.Logger.Debug("templates:%v", templatePages)
	return templatePages
}
