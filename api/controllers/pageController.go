package controllers

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine"

	"github.com/gin-gonic/gin"
	"github.com/kenkoii/Analytics/api/models"
)

type requestForm struct {
	start  string
	end    string
	filter string
}

func HomePageController(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title": "Main website",
		})
}

func ResultPageController(c *gin.Context) {
	// process request here
	ctx, _ := context.WithTimeout(appengine.NewContext(c.Request), time.Second*30)
	start, _ := time.Parse("2006-01-02", c.PostForm("start"))
	end, _ := time.Parse("2006-01-02", c.PostForm("end"))
	filter := c.PostForm("filter")
	table := c.PostForm("table")
	params := gin.H{
		"start":  c.PostForm("start"),
		"end":    c.PostForm("end"),
		"filter": filter,
		"table":  table,
	}

	switch table {
	case "userproperties":
		results, err := models.GetUserProperties(ctx, start, end, filter)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	case "userdailyproperties":
		results, err := models.GetUserDailyProperties(ctx, start, end)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	case "userpurchases":
		results, err := models.GetUserPurchases(ctx, start, end)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	case "tutorials":
		results, err := models.GetTutorials(ctx, start, end)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	case "stages":
		results, err := models.GetStages(ctx, start, end)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	case "gachas":
		results, err := models.GetGachas(ctx, start, end)
		if err != nil {
			LogError(c, err)
		}
		RenderResult(c, results, params)
		break
	}
}

func RenderResult(c *gin.Context, results interface{}, params gin.H) {
	PrintObject(results)
	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title":  "Main website",
			"method": "POST",
			"data":   results,
			"params": params,
		})
}

func PrintObject(obj interface{}) {
	log.Print("===Object===")
	log.Println(obj)
	log.Print("============")
}

func LogError(c *gin.Context, err error) {
	log.Print("===Error===")
	log.Println(err)
	log.Print("===========")
	c.AbortWithError(500, err)
}
