package main

import (
	"net/http"


    "fmt"
    "io"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"


)

type Test struct {
	Age  int
	ScreenWidth int
	ScreenHeight  int
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
    body, _ := io.ReadAll(c.Request.Body)
    fmt.Println(string(body))

    //connect db
    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)
    stmt, err := db.Prepare("INSERT INTO main(age) values(?)")
    checkErr(err)
    res, err := stmt.Exec(12313132);
    checkErr(err)

    //check insert
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
    rows, err := db.Query("SELECT age FROM main WHERE rowid=(?)",id)
    checkErr(err)
    for rows.Next() {
        var age string
        if err := rows.Scan(&age); err != nil {
            checkErr(err)
        }
        fmt.Printf("Age is %s", age)
    }


    c.Redirect(http.StatusFound, "/flow/explain.html?id=321")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


func postFirst(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)
    fmt.Println(string(body))
}


func render(c *gin.Context, template templ.Component, status int) error {
	c.Header("Content-Type", "text/html")
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
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

    r.POST("/flow/age.html", agePost)
    r.POST("/first-ver", postFirst)


	r.Run(":8080")
}
