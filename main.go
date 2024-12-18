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
func getHeatmapPage(c *gin.Context) {
    c.HTML(http.StatusOK, "first-hovers.html", gin.H{})
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

func postFirstSurvey(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)

        verStr := strconv.Itoa(ver)
        id := c.Query("id")
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_`+verStr+`_1_1 = ? ,
                                    survey_`+verStr+`_1_2 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, id );
        checkErr(err)
        fmt.Println(res.RowsAffected());

        c.Redirect(http.StatusFound, "/survey/"+verStr+"-2.html?id=" + id)
    }
    return gin.HandlerFunc(fn)
}

func postSecondSurvey(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")
        q3 := values.Get("quest3")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)


        verStr := strconv.Itoa(ver)
        id := c.Query("id")
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_`+verStr+`_2_1 = ? ,
                                    survey_`+verStr+`_2_2 = ? ,
                                    survey_`+verStr+`_2_3 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, q3, id);
        checkErr(err)
        fmt.Println(res.RowsAffected());

        c.Redirect(http.StatusFound, "/survey/"+verStr+"-3.html?id=" + id)
    }
    return gin.HandlerFunc(fn)
}

func postThirdSurvey(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)


        verStr := strconv.Itoa(ver)
        id := c.Query("id")
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_`+verStr+`_3_1 = ? ,
                                    survey_`+verStr+`_3_2 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, id);
        checkErr(err)
        fmt.Println(res.RowsAffected());

        c.Redirect(http.StatusFound, "/survey/"+verStr+"-4.html?id=" + id)
    }
    return gin.HandlerFunc(fn)
}

func postFourthSurvey(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")
        q3 := values.Get("quest3")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)


        verStr := strconv.Itoa(ver)
        id := c.Query("id")
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_`+verStr+`_4_1 = ? ,
                                    survey_`+verStr+`_4_2 = ? ,
                                    survey_`+verStr+`_4_3 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, q3, id);
        checkErr(err)
        fmt.Println(res.RowsAffected());

        c.Redirect(http.StatusFound, "/survey/"+verStr+"-5.html?id=" + id)
    }
    return gin.HandlerFunc(fn)
}

func postFifthSurvey(ver int) gin.HandlerFunc {
    fn := func(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)


        verStr := strconv.Itoa(ver)
        id := c.Query("id")
        idInt,err := strconv.Atoi(id)
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_`+verStr+`_5_1 = ? ,
                                    survey_`+verStr+`_5_2 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, id);
        checkErr(err)
        fmt.Println(res.RowsAffected());
        
        if ((ver == 1 && idInt % 2 != 0) || (ver == 2 && idInt % 2 == 0)) {
            c.Redirect(http.StatusFound, "/flow/interim.html?id=" + id)
        } else {
            c.Redirect(http.StatusFound, "/survey/6.html?id=" + id)
        }
    }
    return gin.HandlerFunc(fn)
}

func postSixthSurvey(c *gin.Context){
        stringedBody := getBodyString(c)
        //parse stringedBody
        values, err := url.ParseQuery(stringedBody) 
        checkErr(err)

        // Accessing the values
        q1 := values.Get("quest1")
        q2 := values.Get("quest2")

        db, err := sql.Open("sqlite3", "./static/db/thesis.db")
        checkErr(err)


        id := c.Query("id")
        //if not likert number from url
        stmt, err := db.Prepare(`UPDATE main 
                                SET survey_6_1 = ? ,
                                    survey_6_2 = ? 
                                WHERE pid = ? `)
        checkErr(err)
        res, err := stmt.Exec(q1, q2, id);
        checkErr(err)
        fmt.Println(res.RowsAffected());

        c.Redirect(http.StatusFound, "/survey/thank.html?id=" + id)
}
func postSurveyFirst(c *gin.Context) {
    stringedBody := getBodyString(c)
    //parse stringedBody
    values, err := url.ParseQuery(stringedBody) 
    checkErr(err)

    // Accessing the values
    q1 := values.Get("quest1")
    q2 := values.Get("quest2")
    q3 := values.Get("quest3")
    q4 := values.Get("quest4")

    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)

    id := c.Query("id")
    //if not likert number from url
    stmt, err := db.Prepare(`UPDATE main 
                            SET ver_1_1 = ? ,
                                ver_1_2 = ?,
                            WHERE pid = ?`)
    checkErr(err)
    res, err := stmt.Exec(q1, q2, q3, q4, id);
    checkErr(err)
    fmt.Println(res.RowsAffected());

    c.Redirect(http.StatusFound, "/survey/4.html?id=" + id)
}

func postSurvey(c *gin.Context) {
    stringedBody := getBodyString(c)
    answer := getAnswerFromBody(stringedBody)
    fmt.Println("choice: "+answer)


    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)

    pageUrl := c.Param("pageUrl")
    id := c.Query("id")
    //if not likert
    //get questionNumber from url param
    questionNumber := strings.Split(pageUrl, ".")[0]
    stmt, err := db.Prepare("UPDATE main SET survey_"+ questionNumber +" = ? WHERE pid = ?")
    checkErr(err)
    res, err := stmt.Exec(answer, id);
    checkErr(err)
    fmt.Println(res.RowsAffected());


    questionNumberInt,err := strconv.Atoi(questionNumber)
    checkErr(err)
    nextQuestionNumber := questionNumberInt + 1 
    if (nextQuestionNumber < 7) {
        if (questionNumberInt == 2 ){
            c.Redirect(http.StatusFound, "/likert1?id=" + id)
        } else if (questionNumberInt == 3 ){
            c.Redirect(http.StatusFound, "/likert2?id=" + id)
        } else {
            c.Redirect(http.StatusFound, "/survey/"  + strconv.Itoa(nextQuestionNumber) + ".html?id=" + id)
        }
    } else {
        c.Redirect(http.StatusFound, "/survey/thank.html?id=" + id)
    }
}


func getFirstMaxHover(c *gin.Context) {
    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)

    testid := c.Query("id")
    fmt.Println("testid: " + testid)
    testidInt, err := strconv.Atoi(testid)
    checkErr(err)

    //odd find not null
    rows, err := db.Query("SELECT hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1, hover_big_feat_2, hover_big_feat_3, hover_big_feat_4 FROM main WHERE pid=(?) AND age IS NOT NULL",testidInt)
    checkErr(err)
    //even find null
    if (testidInt % 2 == 0) {
        rows,err = db.Query("SELECT hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1, hover_big_feat_2, hover_big_feat_3, hover_big_feat_4 FROM main WHERE pid=(?) AND age IS NULL",testidInt)
        checkErr(err)
    }
    hovers := make(map[string]float32)
    for rows.Next() {
        var hero float32
        var featlist float32
        var benefit float32
        var bigfeat1 float32
        var bigfeat2 float32
        var bigfeat3 float32
        var bigfeat4 float32
        if err := rows.Scan(&hero, &featlist, &benefit, &bigfeat1, &bigfeat2, &bigfeat3, &bigfeat4); err != nil {
            checkErr(err)
        }
        hovers["hero"] = hero
        hovers["featlist"] = featlist
        hovers["benefit"] = benefit
        hovers["bigfeat1"] = bigfeat1
        hovers["bigfeat2"] = bigfeat2
        hovers["bigfeat3"] = bigfeat3
        hovers["bigfeat4"] = bigfeat4

    }
    fmt.Println("sections hover time: ",hovers )

    // Find the maximum hover time
    var maxSection string
    var maxHoverTime float32

    for section, hoverTime := range hovers {
        if hoverTime > maxHoverTime {
            maxHoverTime = hoverTime
            maxSection = section
        }
    }

    fmt.Printf("Max hover time is %f in section: %s\n", maxHoverTime, maxSection)
    c.JSON(http.StatusOK, gin.H{"max_hover_time": maxHoverTime, "section": maxSection})
}

func getSecondMaxHover(c *gin.Context) {
    db, err := sql.Open("sqlite3", "./static/db/thesis.db")
    checkErr(err)

    testid := c.Query("id")
    fmt.Println("testid: " + testid)
    testidInt, err := strconv.Atoi(testid)
    checkErr(err)
    //odd find null
    rows, err := db.Query("SELECT hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1, hover_big_feat_2, hover_big_feat_3, hover_big_feat_4 FROM main WHERE pid=(?) AND age IS NULL",testidInt)
    checkErr(err)
    //even find not null
    if (testidInt % 2 == 0) {
        rows,err = db.Query("SELECT hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1, hover_big_feat_2, hover_big_feat_3, hover_big_feat_4 FROM main WHERE pid=(?) AND age IS NOT NULL",testidInt)
        checkErr(err)
    }
    hovers := make(map[string]float32)
    for rows.Next() {
        var hero float32
        var featlist float32
        var benefit float32
        var bigfeat1 float32
        var bigfeat2 float32
        var bigfeat3 float32
        var bigfeat4 float32
        if err := rows.Scan(&hero, &featlist, &benefit, &bigfeat1, &bigfeat2, &bigfeat3, &bigfeat4); err != nil {
            checkErr(err)
        }
        hovers["hero"] = hero
        hovers["featlist"] = featlist
        hovers["benefit"] = benefit
        hovers["bigfeat1"] = bigfeat1
        hovers["bigfeat2"] = bigfeat2
        hovers["bigfeat3"] = bigfeat3
        hovers["bigfeat4"] = bigfeat4

    }
    fmt.Println("sections hover time: ",hovers )

    // Find the maximum hover time
    var maxSection string
    var maxHoverTime float32

    for section, hoverTime := range hovers {
        if hoverTime > maxHoverTime {
            maxHoverTime = hoverTime
            maxSection = section
        }
    }

    fmt.Printf("Max hover time is %f in section: %s\n", maxHoverTime, maxSection)
    c.JSON(http.StatusOK, gin.H{"max_hover_time": maxHoverTime, "section": maxSection})
}

func getMinHovers(c *gin.Context) {
    db, err := sql.Open("sqlite3", "analyze/thesis.db")
    checkErr(err)

    //odd find not null
    rows, err := db.Query(`SELECT hover_nav_feat, hover_nav_price, hover_nav_login, hover_nav_start,
                                hover_hero_cta, hover_hero_login, hover_small_feat1_pic, hover_small_feat2_pic,
                                hover_small_feat3_pic, hover_headstart, hover_consistency, hover_determination,
                                hover_big_feat1_img, hover_big_feat2_img, hover_big_feat3_img, hover_big_feat4_img,
                                hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1,
                                hover_big_feat_2, hover_big_feat_3, hover_big_feat_4, hover_head_logo,
                                hover_hero_title, hover_sub_title, hover_headstart_desc, hover_consistency_desc,
                                hover_flexible_desc, hover_determination_desc, hover_big_feat1_desc,
                                hover_big_feat2_desc, hover_bigfeat2_cta, hover_big_feat3_desc, hover_bigfeat3_more,
                                hover_big_feat4_desc, hover_bigfeat4_more, hover_ending_title, hover_ending_subtitle,
                                hover_ending_cta_btn, hover_footer_logo, hover_footer_product, hover_footer_company,
                                hover_footer_legal 
                            FROM main 
                            WHERE hover_hero = 1.959
                            ORDER BY top ASC LIMIT 1;`)
    checkErr(err)
    hovers := make(map[string]float32)
    for rows.Next() {
        var navFeat, navPrice, navLogin, navStart float32
        var heroCTA, heroLogin, smallFeat1Pic, smallFeat2Pic, smallFeat3Pic float32
        var headstart, consistency, determination float32
        var bigFeat1Img, bigFeat2Img, bigFeat3Img, bigFeat4Img float32
        var hero, featList, benefitList float32
        var bigFeat1, bigFeat2, bigFeat3, bigFeat4 float32
        var headLogo, heroTitle, subTitle float32
        var headstartDesc, consistencyDesc, flexibleDesc, determinationDesc float32
        var bigFeat1Desc, bigFeat2Desc, bigFeat2CTA, bigFeat3Desc, bigFeat3More float32
        var bigFeat4Desc, bigFeat4More, endingTitle, endingSubtitle float32
        var endingCTABtn, footerLogo, footerProduct, footerCompany, footerLegal float32

        if err := rows.Scan(
            &navFeat, &navPrice, &navLogin, &navStart, &heroCTA, &heroLogin, 
            &smallFeat1Pic, &smallFeat2Pic, &smallFeat3Pic, &headstart, 
            &consistency, &determination, &bigFeat1Img, &bigFeat2Img, 
            &bigFeat3Img, &bigFeat4Img, &hero, &featList, &benefitList, 
            &bigFeat1, &bigFeat2, &bigFeat3, &bigFeat4, &headLogo, 
            &heroTitle, &subTitle, &headstartDesc, &consistencyDesc, 
            &flexibleDesc, &determinationDesc, &bigFeat1Desc, &bigFeat2Desc, 
            &bigFeat2CTA, &bigFeat3Desc, &bigFeat3More, &bigFeat4Desc, 
            &bigFeat4More, &endingTitle, &endingSubtitle, &endingCTABtn, 
            &footerLogo, &footerProduct, &footerCompany, &footerLegal,
        ); 
        err != nil {
            checkErr(err)
        }

        hovers = map[string]float32{
            "nav-feat": navFeat,
            "nav-price": navPrice,
            "nav-login": navLogin,
            "nav-start": navStart,
            "hero-cta": heroCTA,
            "hero-login": heroLogin,
            "small-feat1-pic": smallFeat1Pic,
            "small-feat2-pic": smallFeat2Pic,
            "small-feat3-pic": smallFeat3Pic,
            "headstart": headstart,
            "consistency": consistency,
            "determination": determination,
            "big-feat1-img": bigFeat1Img,
            "big-feat2-img": bigFeat2Img,
            "big-feat3-img": bigFeat3Img,
            "big-feat4-img": bigFeat4Img,
            "hero": hero,
            "feat-list": featList,
            "benefit-list": benefitList,
            "big-feat-1": bigFeat1,
            "big-feat-2": bigFeat2,
            "big-feat-3": bigFeat3,
            "big-feat-4": bigFeat4,
            "head-logo": headLogo,
            "hero-title": heroTitle,
            "sub-title": subTitle,
            "headstart-desc": headstartDesc,
            "consistency-desc": consistencyDesc,
            "flexible-desc": flexibleDesc,
            "determination-desc": determinationDesc,
            "big-feat1-desc": bigFeat1Desc,
            "big-feat2-desc": bigFeat2Desc,
            "big-feat2-cta": bigFeat2CTA,
            "big-feat3-desc": bigFeat3Desc,
            "big-feat3-more": bigFeat3More,
            "big-feat4-desc": bigFeat4Desc,
            "big-feat4-more": bigFeat4More,
            "ending-title": endingTitle,
            "ending-subtitle": endingSubtitle,
            "ending-cta-btn": endingCTABtn,
            "footer-logo": footerLogo,
            "footer-product": footerProduct,
            "footer-company": footerCompany,
            "footer-legal": footerLegal,
        }
    }
    fmt.Println("sections hover time: ",hovers )

    fmt.Printf("Hover times: %s\n", hovers)
    c.JSON(http.StatusOK, gin.H{"Hover time": hovers})
}
func getMaxHovers(c *gin.Context) {
    db, err := sql.Open("sqlite3", "analyze/thesis.db")
    checkErr(err)

    //odd find not null
    rows, err := db.Query(`SELECT hover_nav_feat, hover_nav_price, hover_nav_login, hover_nav_start,
                                hover_hero_cta, hover_hero_login, hover_small_feat1_pic, hover_small_feat2_pic,
                                hover_small_feat3_pic, hover_headstart, hover_consistency, hover_determination,
                                hover_big_feat1_img, hover_big_feat2_img, hover_big_feat3_img, hover_big_feat4_img,
                                hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1,
                                hover_big_feat_2, hover_big_feat_3, hover_big_feat_4, hover_head_logo,
                                hover_hero_title, hover_sub_title, hover_headstart_desc, hover_consistency_desc,
                                hover_flexible_desc, hover_determination_desc, hover_big_feat1_desc,
                                hover_big_feat2_desc, hover_bigfeat2_cta, hover_big_feat3_desc, hover_bigfeat3_more,
                                hover_big_feat4_desc, hover_bigfeat4_more, hover_ending_title, hover_ending_subtitle,
                                hover_ending_cta_btn, hover_footer_logo, hover_footer_product, hover_footer_company,
                                hover_footer_legal 
                            FROM main 
                            WHERE (pid % 2 <> 0 AND version = 1) OR (pid % 2 = 0 AND version = 2)
                            ORDER BY top DESC LIMIT 1 OFFSET 1;`)
    checkErr(err)
    hovers := make(map[string]float32)
    for rows.Next() {
        var navFeat, navPrice, navLogin, navStart float32
        var heroCTA, heroLogin, smallFeat1Pic, smallFeat2Pic, smallFeat3Pic float32
        var headstart, consistency, determination float32
        var bigFeat1Img, bigFeat2Img, bigFeat3Img, bigFeat4Img float32
        var hero, featList, benefitList float32
        var bigFeat1, bigFeat2, bigFeat3, bigFeat4 float32
        var headLogo, heroTitle, subTitle float32
        var headstartDesc, consistencyDesc, flexibleDesc, determinationDesc float32
        var bigFeat1Desc, bigFeat2Desc, bigFeat2CTA, bigFeat3Desc, bigFeat3More float32
        var bigFeat4Desc, bigFeat4More, endingTitle, endingSubtitle float32
        var endingCTABtn, footerLogo, footerProduct, footerCompany, footerLegal float32

        if err := rows.Scan(
            &navFeat, &navPrice, &navLogin, &navStart, &heroCTA, &heroLogin, 
            &smallFeat1Pic, &smallFeat2Pic, &smallFeat3Pic, &headstart, 
            &consistency, &determination, &bigFeat1Img, &bigFeat2Img, 
            &bigFeat3Img, &bigFeat4Img, &hero, &featList, &benefitList, 
            &bigFeat1, &bigFeat2, &bigFeat3, &bigFeat4, &headLogo, 
            &heroTitle, &subTitle, &headstartDesc, &consistencyDesc, 
            &flexibleDesc, &determinationDesc, &bigFeat1Desc, &bigFeat2Desc, 
            &bigFeat2CTA, &bigFeat3Desc, &bigFeat3More, &bigFeat4Desc, 
            &bigFeat4More, &endingTitle, &endingSubtitle, &endingCTABtn, 
            &footerLogo, &footerProduct, &footerCompany, &footerLegal,
        ); 
        err != nil {
            checkErr(err)
        }

        hovers = map[string]float32{
            "nav-feat": navFeat,
            "nav-price": navPrice,
            "nav-login": navLogin,
            "nav-start": navStart,
            "hero-cta": heroCTA,
            "hero-login": heroLogin,
            "small-feat1-pic": smallFeat1Pic,
            "small-feat2-pic": smallFeat2Pic,
            "small-feat3-pic": smallFeat3Pic,
            "headstart": headstart,
            "consistency": consistency,
            "determination": determination,
            "big-feat1-img": bigFeat1Img,
            "big-feat2-img": bigFeat2Img,
            "big-feat3-img": bigFeat3Img,
            "big-feat4-img": bigFeat4Img,
            "hero": hero,
            "feat-list": featList,
            "benefit-list": benefitList,
            "big-feat-1": bigFeat1,
            "big-feat-2": bigFeat2,
            "big-feat-3": bigFeat3,
            "big-feat-4": bigFeat4,
            "head-logo": headLogo,
            "hero-title": heroTitle,
            "sub-title": subTitle,
            "headstart-desc": headstartDesc,
            "consistency-desc": consistencyDesc,
            "flexible-desc": flexibleDesc,
            "determination-desc": determinationDesc,
            "big-feat1-desc": bigFeat1Desc,
            "big-feat2-desc": bigFeat2Desc,
            "big-feat2-cta": bigFeat2CTA,
            "big-feat3-desc": bigFeat3Desc,
            "big-feat3-more": bigFeat3More,
            "big-feat4-desc": bigFeat4Desc,
            "big-feat4-more": bigFeat4More,
            "ending-title": endingTitle,
            "ending-subtitle": endingSubtitle,
            "ending-cta-btn": endingCTABtn,
            "footer-logo": footerLogo,
            "footer-product": footerProduct,
            "footer-company": footerCompany,
            "footer-legal": footerLegal,
        }
    }
    fmt.Println("sections hover time: ",hovers )

    fmt.Printf("Hover times: %s\n", hovers)
    c.JSON(http.StatusOK, gin.H{"Hover time": hovers})
}

func getFirstHovers(c *gin.Context) {
    db, err := sql.Open("sqlite3", "analyze/thesis.db")
    checkErr(err)

    testid := c.Query("id")
    fmt.Println("testid: " + testid)
    testidInt, err := strconv.Atoi(testid)
    checkErr(err)

    //odd find not null
    rows, err := db.Query(`SELECT hover_nav_feat, hover_nav_price, hover_nav_login, hover_nav_start,
                                hover_hero_cta, hover_hero_login, hover_small_feat1_pic, hover_small_feat2_pic,
                                hover_small_feat3_pic, hover_headstart, hover_consistency, hover_determination,
                                hover_big_feat1_img, hover_big_feat2_img, hover_big_feat3_img, hover_big_feat4_img,
                                hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1,
                                hover_big_feat_2, hover_big_feat_3, hover_big_feat_4, hover_head_logo,
                                hover_hero_title, hover_sub_title, hover_headstart_desc, hover_consistency_desc,
                                hover_flexible_desc, hover_determination_desc, hover_big_feat1_desc,
                                hover_big_feat2_desc, hover_bigfeat2_cta, hover_big_feat3_desc, hover_bigfeat3_more,
                                hover_big_feat4_desc, hover_bigfeat4_more, hover_ending_title, hover_ending_subtitle,
                                hover_ending_cta_btn, hover_footer_logo, hover_footer_product, hover_footer_company,
                                hover_footer_legal 
                            FROM main 
                            WHERE pid = (?) ORDER BY ROWID ASC LIMIT 1;`,testidInt)
    checkErr(err)
    //even find null
    //TODO: copy aboved query to this SELECT
    if (testidInt % 2 == 0) {
        rows,err = db.Query("SELECT hover_hero, hover_feat_list, hover_benefit_list, hover_big_feat_1, hover_big_feat_2, hover_big_feat_3, hover_big_feat_4 FROM main WHERE pid=(?) AND age IS NULL",testidInt)
        checkErr(err)
    }
    hovers := make(map[string]float32)
    for rows.Next() {
        var navFeat, navPrice, navLogin, navStart float32
        var heroCTA, heroLogin, smallFeat1Pic, smallFeat2Pic, smallFeat3Pic float32
        var headstart, consistency, determination float32
        var bigFeat1Img, bigFeat2Img, bigFeat3Img, bigFeat4Img float32
        var hero, featList, benefitList float32
        var bigFeat1, bigFeat2, bigFeat3, bigFeat4 float32
        var headLogo, heroTitle, subTitle float32
        var headstartDesc, consistencyDesc, flexibleDesc, determinationDesc float32
        var bigFeat1Desc, bigFeat2Desc, bigFeat2CTA, bigFeat3Desc, bigFeat3More float32
        var bigFeat4Desc, bigFeat4More, endingTitle, endingSubtitle float32
        var endingCTABtn, footerLogo, footerProduct, footerCompany, footerLegal float32

        if err := rows.Scan(
            &navFeat, &navPrice, &navLogin, &navStart, &heroCTA, &heroLogin, 
            &smallFeat1Pic, &smallFeat2Pic, &smallFeat3Pic, &headstart, 
            &consistency, &determination, &bigFeat1Img, &bigFeat2Img, 
            &bigFeat3Img, &bigFeat4Img, &hero, &featList, &benefitList, 
            &bigFeat1, &bigFeat2, &bigFeat3, &bigFeat4, &headLogo, 
            &heroTitle, &subTitle, &headstartDesc, &consistencyDesc, 
            &flexibleDesc, &determinationDesc, &bigFeat1Desc, &bigFeat2Desc, 
            &bigFeat2CTA, &bigFeat3Desc, &bigFeat3More, &bigFeat4Desc, 
            &bigFeat4More, &endingTitle, &endingSubtitle, &endingCTABtn, 
            &footerLogo, &footerProduct, &footerCompany, &footerLegal,
        ); 
        err != nil {
            checkErr(err)
        }

        hovers = map[string]float32{
            "nav-feat": navFeat,
            "nav-price": navPrice,
            "nav-login": navLogin,
            "nav-start": navStart,
            "hero-cta": heroCTA,
            "hero-login": heroLogin,
            "small-feat1-pic": smallFeat1Pic,
            "small-feat2-pic": smallFeat2Pic,
            "small-feat3-pic": smallFeat3Pic,
            "headstart": headstart,
            "consistency": consistency,
            "determination": determination,
            "big-feat1-img": bigFeat1Img,
            "big-feat2-img": bigFeat2Img,
            "big-feat3-img": bigFeat3Img,
            "big-feat4-img": bigFeat4Img,
            "hero": hero,
            "feat-list": featList,
            "benefit-list": benefitList,
            "big-feat-1": bigFeat1,
            "big-feat-2": bigFeat2,
            "big-feat-3": bigFeat3,
            "big-feat-4": bigFeat4,
            "head-logo": headLogo,
            "hero-title": heroTitle,
            "sub-title": subTitle,
            "headstart-desc": headstartDesc,
            "consistency-desc": consistencyDesc,
            "flexible-desc": flexibleDesc,
            "determination-desc": determinationDesc,
            "big-feat1-desc": bigFeat1Desc,
            "big-feat2-desc": bigFeat2Desc,
            "big-feat2-cta": bigFeat2CTA,
            "big-feat3-desc": bigFeat3Desc,
            "big-feat3-more": bigFeat3More,
            "big-feat4-desc": bigFeat4Desc,
            "big-feat4-more": bigFeat4More,
            "ending-title": endingTitle,
            "ending-subtitle": endingSubtitle,
            "ending-cta-btn": endingCTABtn,
            "footer-logo": footerLogo,
            "footer-product": footerProduct,
            "footer-company": footerCompany,
            "footer-legal": footerLegal,
        }
    }
    fmt.Println("sections hover time: ",hovers )

    fmt.Printf("Hover times: %s\n", hovers)
    c.JSON(http.StatusOK, gin.H{"Hover time": hovers})
}


func getAVGFirstHovers(c *gin.Context) {
    db, err := sql.Open("sqlite3", "analyze/thesis.db")
    checkErr(err)

    //odd find not null
    rows, err := db.Query(`SELECT AVG(hover_nav_feat) as hover_nav_feat, AVG(hover_nav_price) as hover_nav_price, AVG(hover_nav_login) as hover_nav_login, AVG(hover_nav_start) as hover_nav_start,
                                AVG(hover_hero_cta) as hover_hero_cta, AVG(hover_hero_login) as hover_hero_login, AVG(hover_small_feat1_pic) as hover_small_feat1_pic, AVG(hover_small_feat2_pic) as hover_small_feat2_pic,
                                AVG(hover_small_feat3_pic) hover_small_feat3_pic, AVG(hover_headstart) as hover_headstart, AVG(hover_consistency) as hover_consistency, AVG(hover_determination) as hover_determination,
                                AVG(hover_big_feat1_img) as hover_big_feat1_img, AVG(hover_big_feat2_img) as hover_big_feat2_img, AVG(hover_big_feat3_img) as hover_big_feat3_img, AVG(hover_big_feat4_img) as hover_big_feat4_img,
                                AVG(hover_hero) as hover_hero, AVG(hover_feat_list) as hover_feat_list, AVG(hover_benefit_list) as hover_benefit_list, AVG(hover_big_feat_1) as hover_big_feat_1,
                                AVG(hover_big_feat_2) as hover_big_feat_2, AVG(hover_big_feat_3) as hover_big_feat_3, AVG(hover_big_feat_4) as hover_big_feat_4, AVG(hover_head_logo) as hover_head_logo,
                                AVG(hover_hero_title) as hover_hero_title, AVG(hover_sub_title) as hover_sub_title, AVG(hover_headstart_desc) as hover_headstart_desc, AVG(hover_consistency_desc) as hover_consistency_desc,
                                AVG(hover_flexible_desc) as hover_flexible_desc, AVG(hover_determination_desc) as hover_determination_desc, AVG(hover_big_feat1_desc) as hover_big_feat1_desc,
                                AVG(hover_big_feat2_desc) as hover_big_feat2_desc, AVG(hover_bigfeat2_cta) as hover_bigfeat2_cta, AVG(hover_big_feat3_desc) as hover_big_feat3_desc, AVG(hover_bigfeat3_more) as hover_bigfeat3_more,
                                AVG(hover_big_feat4_desc) as hover_big_feat4_desc, AVG(hover_bigfeat4_more) as hover_bigfeat4_more, AVG(hover_ending_title) as hover_ending_title, AVG(hover_ending_subtitle) as hover_ending_subtitle,
                                AVG(hover_ending_cta_btn) as hover_ending_cta_btn, AVG(hover_footer_logo) as hover_footer_logo, AVG(hover_footer_product) as hover_footer_product, AVG(hover_footer_company) as hover_footer_company,
                                AVG(hover_footer_legal) as hover_footer_legal
                            FROM main`)
    checkErr(err)
    hovers := make(map[string]float32)
    for rows.Next() {
        var navFeat, navPrice, navLogin, navStart float32
        var heroCTA, heroLogin, smallFeat1Pic, smallFeat2Pic, smallFeat3Pic float32
        var headstart, consistency, determination float32
        var bigFeat1Img, bigFeat2Img, bigFeat3Img, bigFeat4Img float32
        var hero, featList, benefitList float32
        var bigFeat1, bigFeat2, bigFeat3, bigFeat4 float32
        var headLogo, heroTitle, subTitle float32
        var headstartDesc, consistencyDesc, flexibleDesc, determinationDesc float32
        var bigFeat1Desc, bigFeat2Desc, bigFeat2CTA, bigFeat3Desc, bigFeat3More float32
        var bigFeat4Desc, bigFeat4More, endingTitle, endingSubtitle float32
        var endingCTABtn, footerLogo, footerProduct, footerCompany, footerLegal float32

        if err := rows.Scan(
            &navFeat, &navPrice, &navLogin, &navStart, &heroCTA, &heroLogin, 
            &smallFeat1Pic, &smallFeat2Pic, &smallFeat3Pic, &headstart, 
            &consistency, &determination, &bigFeat1Img, &bigFeat2Img, 
            &bigFeat3Img, &bigFeat4Img, &hero, &featList, &benefitList, 
            &bigFeat1, &bigFeat2, &bigFeat3, &bigFeat4, &headLogo, 
            &heroTitle, &subTitle, &headstartDesc, &consistencyDesc, 
            &flexibleDesc, &determinationDesc, &bigFeat1Desc, &bigFeat2Desc, 
            &bigFeat2CTA, &bigFeat3Desc, &bigFeat3More, &bigFeat4Desc, 
            &bigFeat4More, &endingTitle, &endingSubtitle, &endingCTABtn, 
            &footerLogo, &footerProduct, &footerCompany, &footerLegal,
        ); 
        err != nil {
            checkErr(err)
        }

        hovers = map[string]float32{
            "nav-feat": navFeat,
            "nav-price": navPrice,
            "nav-login": navLogin,
            "nav-start": navStart,
            "hero-cta": heroCTA,
            "hero-login": heroLogin,
            "small-feat1-pic": smallFeat1Pic,
            "small-feat2-pic": smallFeat2Pic,
            "small-feat3-pic": smallFeat3Pic,
            "headstart": headstart,
            "consistency": consistency,
            "determination": determination,
            "big-feat1-img": bigFeat1Img,
            "big-feat2-img": bigFeat2Img,
            "big-feat3-img": bigFeat3Img,
            "big-feat4-img": bigFeat4Img,
            "hero": hero,
            "feat-list": featList,
            "benefit-list": benefitList,
            "big-feat-1": bigFeat1,
            "big-feat-2": bigFeat2,
            "big-feat-3": bigFeat3,
            "big-feat-4": bigFeat4,
            "head-logo": headLogo,
            "hero-title": heroTitle,
            "sub-title": subTitle,
            "headstart-desc": headstartDesc,
            "consistency-desc": consistencyDesc,
            "flexible-desc": flexibleDesc,
            "determination-desc": determinationDesc,
            "big-feat1-desc": bigFeat1Desc,
            "big-feat2-desc": bigFeat2Desc,
            "big-feat2-cta": bigFeat2CTA,
            "big-feat3-desc": bigFeat3Desc,
            "big-feat3-more": bigFeat3More,
            "big-feat4-desc": bigFeat4Desc,
            "big-feat4-more": bigFeat4More,
            "ending-title": endingTitle,
            "ending-subtitle": endingSubtitle,
            "ending-cta-btn": endingCTABtn,
            "footer-logo": footerLogo,
            "footer-product": footerProduct,
            "footer-company": footerCompany,
            "footer-legal": footerLegal,
        }
    }
    fmt.Println("sections hover time: ",hovers )

    fmt.Printf("Hover times: %s\n", hovers)
    c.JSON(http.StatusOK, gin.H{"Hover time": hovers})
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
    r.Static("/analyze", "./analyze")
    //Gin can only load one of this function
    r.LoadHTMLFiles("static/index.html", "static/second-ver.html", "static/flow/start.html", "analyze/first-hovers.html")


    r.GET("/", getFlow)
    r.GET("/first-ver", getFirst)
    r.GET("/second-ver", getSecond)

    r.GET("/first-session-time", getFirstMaxHover)
    r.GET("/second-session-time", getSecondMaxHover)

    r.GET("/heatmap-page", getHeatmapPage)
    r.GET("/first-hovers", getFirstHovers)
    r.GET("/avg-first-hovers", getAVGFirstHovers)
    r.GET("/max-hovers", getMaxHovers)
    r.GET("/min-hovers", getMinHovers)



    r.POST("/flow/age.html", agePost)
    r.POST("/first-ver", postLanding(1))
    r.POST("/second-ver", postLanding(2))

    r.POST("/survey/1-1.html", postFirstSurvey(1))
    r.POST("/survey/2-1.html", postFirstSurvey(2))
    r.POST("/survey/1-2.html", postSecondSurvey(1))
    r.POST("/survey/2-2.html", postSecondSurvey(2))
    r.POST("/survey/1-3.html", postThirdSurvey(1))
    r.POST("/survey/2-3.html", postThirdSurvey(2))
    r.POST("/survey/1-4.html", postFourthSurvey(1))
    r.POST("/survey/2-4.html", postFourthSurvey(2))
    r.POST("/survey/1-5.html", postFifthSurvey(1))
    r.POST("/survey/2-5.html", postFifthSurvey(2))
    r.POST("/survey/6.html", postSixthSurvey)


    r.Run(":8080")
}

