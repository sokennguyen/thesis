package main

import (
	"net/http"
	"net/url"


    "fmt"
    "io"
    "strconv" 

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

    "encoding/json"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"


)

var count int = 0;

type Test struct {
	Age  int
	ScreenWidth int
	ScreenHeight  int
}

type AgeScreenData struct {
    Age int
	ScreenWidth int
	ScreenHeight  int
}

type InteractionData struct {
    Hovers map[string]float64 `json:"hovers"`
    Clicks map[string]float64 `json:"clicks"`
}


func getFirst(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getFlow(c *gin.Context) {
    c.HTML(http.StatusOK, "start.html", gin.H{})
}

func postFirst(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    //decode the JSON body
    var interaction InteractionData
    err := json.Unmarshal(body, &interaction )
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing JSON"})
		return
    }

    //combine hovers and clicks into a single map for SQL update
    combinedData := make(map[string]float64)
	for key, value := range interaction.Clicks {
		combinedData[key] = value
	}
    for key, value := range interaction.Hovers {
		combinedData[key] = value
	}
    fmt.Println(combinedData)

    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)
    stmt, err := db.Prepare(`
            UPDATE main 
            SET click_nav_feat = ?,
                click_nav_price = ?,
                click_nav_login = ?,
                click_nav_start = ?,
                click_hero_cta = ?,
                click_hero_login = ?,
                click_small_feat1_pic = ?,
                click_small_feat2_pic = ?,
                click_small_feat3_pic = ?,
                click_headstart = ?,
                click_consistency = ?,
                click_determination = ?,
                click_big_feat1_img = ?,
                click_big_feat2_img = ?,
                click_big_feat3_img = ?,
                click_big_feat4_img = ?,
                click_ending_cta_btn = ?,
                
                top = ?,
                hover_hero = ?,
                hover_feat_list = ?,
                hover_benefit_list = ?,
                hover_big_feat_1 = ?,
                hover_big_feat_2 = ?,
                hover_big_feat_3 = ?,
                hover_big_feat_4 = ?,
                hover_head_logo = ?,
                hover_hero_title = ?,
                hover_sub_title = ?,
                hover_headstart_desc = ?,
                hover_consistency_desc = ?,
                hover_flexible_desc = ?,
                hover_determination_desc = ?,
                hover_big_feat1_desc = ?,
                hover_big_feat2_desc = ?,
                hover_big_feat3_desc = ?,
                hover_big_feat4_desc = ?,
                hover_ending_title = ?,
                hover_ending_subtitle = ?,
                hover_ending_cta_btn = ?,
                hover_footer_logo = ?,
                hover_footer_product = ?,
                hover_footer_company = ?,
                hover_footer_legal = ? 
            WHERE rowid = ?
    `)
    checkErr(err)

    rowid := c.Query("id")
    fmt.Println("rowid: " + rowid)

    if rowid == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rowid parameter"})
        return
    }


    data := []interface{}{
        combinedData["nav-feat"],
        combinedData["nav-price"],
        combinedData["nav-login"],
        combinedData["nav-start"],
        combinedData["hero-cta"],
        combinedData["hero-login"],
        combinedData["small-feat1-pic"],
        combinedData["small-feat2-pic"],
        combinedData["small-feat3-pic"],
        combinedData["headstart"],
        combinedData["consistency"],
        combinedData["determination"],
        combinedData["big-feat1-img"],
        combinedData["big-feat2-img"],
        combinedData["big-feat3-img"],
        combinedData["big-feat4-img"],
        combinedData["ending-cta-btn"],

        combinedData["top"],
        combinedData["hero"],
        combinedData["feat-list"],
        combinedData["benefit-list"],
        combinedData["big-feat-1"],
        combinedData["big-feat-2"],
        combinedData["big-feat-3"],
        combinedData["big-feat-4"],
        combinedData["head-logo"],
        combinedData["hero-title"],
        combinedData["sub-title"],
        combinedData["headstart-desc"],
        combinedData["consistency-desc"],
        combinedData["flexible-desc"],
        combinedData["determination-desc"],
        combinedData["big-feat1-desc"],
        combinedData["big-feat2-desc"],
        combinedData["big-feat3-desc"],
        combinedData["big-feat4-desc"],
        combinedData["ending-title"],
        combinedData["ending-subtitle"],
        combinedData["ending-cta-btn"],
        combinedData["footer-logo"],
        combinedData["footer-product"],
        combinedData["footer-company"],
        combinedData["footer-legal"],

        rowid, // The user ID to identify the row to update
    }

    _,err = stmt.Exec(data...)
    checkErr(err)
}

func agePost(c *gin.Context) {
    //TODO: gen new ID here
    //TODO: get age from response here
    body, _ := io.ReadAll(c.Request.Body)
    fmt.Println(string(body)) //this value is byte
    stringedBody := string(body)

    //parse stringedBody
    values, err := url.ParseQuery(stringedBody) 
    checkErr(err)

    // Accessing the values
    age := values.Get("age")
    swidth := values.Get("swidth")
    sheight := values.Get("sheight")


    //connect and insert 
    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)
    stmt, err := db.Prepare("INSERT INTO main(age, screen_width, screen_height) values(?,?,?)")
    checkErr(err)
    res, err := stmt.Exec( age, swidth, sheight);
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

    count = count+1;
    c.Redirect(http.StatusFound, "/flow/explain.html?id="+strconv.Itoa(count))
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
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

