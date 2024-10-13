package main

import (
	"net/http"


	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Count struct {
	Count int
}

func render(c *gin.Context, template templ.Component, status int) error {
	c.Header("Content-Type", "text/html")
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}


func getHtml(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {

    r := gin.Default()

    r.Static("/assets", "./static/assets")
    r.StaticFile("/normalize.css", "./static/normalize.css")
    r.StaticFile("/style.css", "./static/style.css")
    r.LoadHTMLFiles("static/index.html")
	r.GET("/", getHtml)

	r.Run(":8080")
}
