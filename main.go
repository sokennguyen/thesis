package main

import (
	"net/http"
	"net/url"

	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"database/sql"
	"encoding/json"

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
func getSecond(c *gin.Context) {
    c.HTML(http.StatusOK, "second-ver.html", gin.H{})
}

func getFlow(c *gin.Context) {
    c.HTML(http.StatusOK, "start.html", gin.H{})
}

//gin.HandlerFunc return *gin.Context
func postLanding(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
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
            if combinedData[key] != 0 {
                combinedData["hover-"+key] = value
            } else {
                combinedData[key] = value 
            }
        }
        fmt.Println(combinedData)

        testid := c.Query("id")
        fmt.Println("testid: " + testid)
        fmt.Println("version: " + strconv.Itoa(ver))
        testidInt, err := strconv.Atoi(testid)
        checkErr(err)

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)
         
        //check if 1st phase is done or not
        //if version is not empty then first post of 2nd phase will INSERT a new row
        //if version is empty then update everything
        rows, err := db.Query("SELECT age, version FROM main WHERE pid = (?)",testid)
        checkErr(err)

        var age sql.NullInt16
        var version sql.NullInt16
        rowsCount := 0
        for rows.Next() {
            if err := rows.Scan(&age, &version ); err != nil {
                checkErr(err)
            }
            rowsCount++
            fmt.Println(age.Valid)
            fmt.Println(version.Valid)
        }
        if age.Valid  && !(version.Valid) && rowsCount == 1 {
                setVerStmt, err := db.Prepare("UPDATE main SET version = ? WHERE pid = ?")
                checkErr(err)
                _,err = setVerStmt.Exec(ver, testid)
                checkErr(err)
                fmt.Println("update ended")
                insertStmt, err := db.Prepare("INSERT INTO main(pid) VALUES (?)")
                checkErr(err)
                _,err = insertStmt.Exec(testidInt)
                checkErr(err)
                fmt.Println("insert ended")
                fmt.Println("Second phase row inserted")
                setVerStmt, err = db.Prepare("UPDATE main SET version = ? WHERE pid = ? AND version is null")
                checkErr(err)
                if (ver==1) {
                    _,err = setVerStmt.Exec(2,testidInt)
                    checkErr(err)
                } else if (ver == 2) {
                    _,err = setVerStmt.Exec(1,testidInt)
                    checkErr(err)
                }
                fmt.Println("Second phase row's version update")
        }

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
                    click_bigfeat2_cta = ?,
                    click_big_feat3_img = ?,
                    click_bigfeat3_more = ?,
                    click_big_feat4_img = ?,
                    click_bigfeat4_more = ?,
                    click_ending_cta_btn = ?,
                    
                    hover_nav_feat = ?,
                    hover_nav_price = ?,
                    hover_nav_login = ?,
                    hover_nav_start = ?,
                    hover_hero_cta = ?,
                    hover_hero_login = ?,
                    hover_small_feat1_pic = ?,
                    hover_small_feat2_pic = ?,
                    hover_small_feat3_pic = ?,
                    hover_headstart = ?,
                    hover_consistency = ?,
                    hover_determination = ?,
                    hover_big_feat1_img = ?,
                    hover_big_feat2_img = ?,
                    hover_bigfeat2_cta = ?,
                    hover_big_feat3_img = ?,
                    hover_bigfeat3_more = ?,
                    hover_big_feat4_img = ?,
                    hover_bigfeat4_more = ?,
                    hover_ending_cta_btn = ?,

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
                    hover_bigfeat2_cta = ?,
                    hover_big_feat3_desc = ?,
                    hover_bigfeat3_more = ?,
                    hover_big_feat4_desc = ?,
                    hover_bigfeat4_more = ?,
                    hover_ending_title = ?,
                    hover_ending_subtitle = ?,
                    hover_ending_cta_btn = ?,
                    hover_footer_logo = ?,
                    hover_footer_product = ?,
                    hover_footer_company = ?,
                    hover_footer_legal = ? 
                WHERE pid = ? AND version = ?
        `)
        checkErr(err)


        if testid == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing testid parameter"})
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
            combinedData["big-feat2-cta"],
            combinedData["big-feat3-img"],
            combinedData["big-feat3-more"],
            combinedData["big-feat4-img"],
            combinedData["big-feat4-more"],
            combinedData["ending-cta-btn"],

            combinedData["hover-nav-feat"],
            combinedData["hover-nav-price"],
            combinedData["hover-nav-login"],
            combinedData["hover-nav-start"],
            combinedData["hover-hero-cta"],
            combinedData["hover-hero-login"],
            combinedData["hover-small-feat1-pic"],
            combinedData["hover-small-feat2-pic"],
            combinedData["hover-small-feat3-pic"],
            combinedData["hover-headstart"],
            combinedData["hover-consistency"],
            combinedData["hover-determination"],
            combinedData["hover-big-feat1-img"],
            combinedData["hover-big-feat2-img"],
            combinedData["hover-big-feat2-cta"],
            combinedData["hover-big-feat3-img"],
            combinedData["hover-big-feat3-more"],
            combinedData["hover-big-feat4-img"],
            combinedData["hover-big-feat4-more"],
            combinedData["hover-ending-cta-btn"],

            combinedData["top"],
            combinedData["hover-hero"],
            combinedData["hover-feat-list"],
            combinedData["hover-benefit-list"],
            combinedData["hover-big-feat-1"],
            combinedData["hover-big-feat-2"],
            combinedData["hover-big-feat-3"],
            combinedData["hover-big-feat-4"],
            combinedData["hover-head-logo"],
            combinedData["hover-hero-title"],
            combinedData["hover-sub-title"],
            combinedData["hover-headstart-desc"],
            combinedData["hover-consistency-desc"],
            combinedData["hover-flexible-desc"],
            combinedData["hover-determination-desc"],
            combinedData["hover-big-feat1-desc"],
            combinedData["hover-big-feat2-desc"],
            combinedData["hover-big-feat2-cta"],
            combinedData["hover-big-feat3-desc"],
            combinedData["hover-big-feat3-more"],
            combinedData["hover-big-feat4-desc"],
            combinedData["hover-big-feat4-more"],
            combinedData["hover-ending-title"],
            combinedData["hover-ending-subtitle"],
            combinedData["hover-ending-cta-btn"],
            combinedData["hover-footer-logo"],
            combinedData["hover-footer-product"],
            combinedData["hover-footer-company"],
            combinedData["hover-footer-legal"],

            testid, // The user ID to identify the row to update
            ver,
        }
        _,err = stmt.Exec(data...)
        checkErr(err)
    }
    return gin.HandlerFunc(fn)
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
    stmt, err := db.Prepare("INSERT INTO main(pid, age, screen_width, screen_height) values(?,?,?,?)")
    checkErr(err)
    res, err := stmt.Exec(count+1, age, swidth, sheight);
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

func getBodyString(c *gin.Context) string {
    body, _ := io.ReadAll(c.Request.Body)
    return string(body)//this value is byte so it needs stringified
    
}
func getAnswerFromBody(body string) string {
    //parse stringed body
    values, err := url.ParseQuery(body) 
    checkErr(err)

    // Accessing the values
    return values.Get("answer")
}

func postSurvey(c *gin.Context) {
    stringedBody := getBodyString(c)
    answer := getAnswerFromBody(stringedBody)
    fmt.Println("choice: "+answer)


    //connect and insert 
    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)

    //get questionNumber from url param
    pageUrl := c.Param("pageUrl")
    questionNumber := strings.Split(pageUrl, ".")[0]
    fmt.Println("questionNumber: "+questionNumber)

    id := c.Query("id")
    stmt, err := db.Prepare("UPDATE main SET survey_"+ questionNumber +" = ? WHERE pid = ?")
    checkErr(err)
    res, err := stmt.Exec(answer, id);
    checkErr(err)
    fmt.Println(res.RowsAffected());


    questionNumberInt,err := strconv.Atoi(questionNumber)
    checkErr(err)
    nextQuestionNumber := questionNumberInt + 1 
    if (nextQuestionNumber < 7) {
        c.Redirect(http.StatusFound, "/survey/"  + strconv.Itoa(nextQuestionNumber) + ".html?id=" + id)
    } else {
        c.Redirect(http.StatusFound, "/survey/thank.html?id=" + id)
    }
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
    r.LoadHTMLFiles("static/index.html", "static/second-ver.html", "static/flow/start.html")


	r.GET("/", getFlow)
	r.GET("/first-ver", getFirst)
	r.GET("/second-ver", getSecond)

    r.POST("/flow/age.html", agePost)
    r.POST("/first-ver", postLanding(1))
    r.POST("/second-ver", postLanding(2))

    r.POST("/survey/:pageUrl", postSurvey)


	r.Run(":8080")
}

