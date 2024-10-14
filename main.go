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


func getFirst(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getFlow(c *gin.Context) {
    c.HTML(http.StatusOK, "start.html", gin.H{})
}


func main() {

    r := gin.Default()

    //load statics
    r.Static("/assets", "./static/assets")
    r.Static("/flow", "./static/flow")
    r.Static("/survey", "./static/survey")
    r.StaticFile("/normalize.css", "./static/normalize.css")
    r.StaticFile("/style.css", "./static/style.css")
    r.StaticFile("/flow.css", "./static/flow.css")
    //Gin can only load one of this function
    r.LoadHTMLFiles("static/index.html", "static/flow/start.html")


	r.GET("/", getFlow)
	r.GET("/first-ver", getFirst)


	r.Run(":8080")
}
