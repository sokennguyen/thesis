package main

import (
	"net/http"
	"net/url"

    "fmt"
    "io"

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

func agePost(c *gin.Context) {
    //TODO: gen new ID here
    //TODO: get age from response here
    c.Redirect(http.StatusFound, "/flow/explain.html?id=321")
}

func getFlowWithReferer(c *gin.Context) {
    referer := c.Request.Referer() // Get the Referer header
    refererUrl,err := url.Parse(referer)
    // Get the parameters
    queryParams := refererURL.Query()
    id := queryParams.Get("id")

    requestedURL := c.Request.URL.String()
    finalUrlWithId := requestedURL + "?id=" + id
    c.Redirect(http.StatusFound, "/flow/explain.html?id=321")
}

func post(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)
    fmt.Println(string(body))
}

func main() {

    r := gin.Default()

    //load statics dirs
    r.Static("/assets", "./static/assets")
    r.Static("/flow", "./static/flow")
    r.Static("/survey", "./static/survey")
    r.Static("/css", "./static/css")
    r.Static("/static", "./static")
    //Gin can only load one of this function
    r.LoadHTMLFiles("static/index.html", "static/flow/start.html")


	r.GET("/", getFlow)
	r.GET("/first-ver", getFirst)
    r.GET("/flow/interim.html", getFlowWithReferer)

    r.POST("/flow/age.html", agePost)


	r.Run(":8080")
}
