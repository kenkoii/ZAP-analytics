package handlers

import (
	"encoding/json"
	x "log"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gin-gonic/gin"
	"github.com/kenkoii/Analytics/api/models"
)

// Handler handles the '/' route
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FreCre Analytics Online"))
}

func Greetings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func PostAnalytics(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	var count = 0
	var analyticsRequests []models.AnalyticsRequest
	if c.BindJSON(&analyticsRequests) == nil {
		for i := 0; i < len(analyticsRequests); i++ {
			switch analyticsRequests[i].ClassName {
			case "UserProperty":
				{
					var userProperty models.UserProperty
					if err := json.Unmarshal(analyticsRequests[i].Params, &userProperty); err != nil {
						LogErrorGin(c, err)
						return
					}
					_, err := models.NewUserProperty(ctx, userProperty)
					if err != nil {
						LogErrorGin(c, err)
						return
					}
					// c.JSON(http.StatusOK, userProperty)
					count++
				}
			case "UserPurchase":
				{
					var userPurchase models.UserPurchase
					if err := json.Unmarshal(analyticsRequests[i].Params, &userPurchase); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewUserPurchase(ctx, userPurchase)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			case "UserDailyProperty":
				{
					var userDailyProperty models.UserDailyProperty
					if err := json.Unmarshal(analyticsRequests[i].Params, &userDailyProperty); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewUserDailyProperty(ctx, userDailyProperty)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			case "Tutorial":
				{
					var tutorial models.Tutorial
					if err := json.Unmarshal(analyticsRequests[i].Params, &tutorial); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewTutorial(ctx, tutorial)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			case "Stage":
				{
					var stage models.Stage
					if err := json.Unmarshal(analyticsRequests[i].Params, &stage); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewStage(ctx, stage)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			case "Quest":
				{
					var quest models.Quest
					if err := json.Unmarshal(analyticsRequests[i].Params, &quest); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewQuest(ctx, quest)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			case "Event":
				{
					var event models.Event
					if err := json.Unmarshal(analyticsRequests[i].Params, &event); err != nil {
						LogErrorGin(c, err)
						return
					}

					_, err := models.NewEvent(ctx, event)
					if err != nil {
						// http.Error(w, err.Error(), http.StatusInternalServerError)
						LogErrorGin(c, err)
						return
					}
					count++
				}
			}
		}
		c.String(http.StatusOK, "1")
	} else {
		c.String(500, "Cant Make JSON")
	}
}

func GetUserProperties(c *gin.Context) {
	ctx, _ := context.WithTimeout(appengine.NewContext(c.Request), time.Second*30)
	if err := CheckTime(c); err != "" {
		c.String(http.StatusInternalServerError, err)
	}

	start, err := ParseTime(c.Query("start"))
	if err != nil {
		c.Error(err)
	}

	end, err := ParseTime(c.Query("end"))
	if err != nil {
		c.Error(err)
	}

	filter := c.Query("filter")
	if filter == "" {
		c.String(http.StatusInternalServerError, "Request Error: Filter param missing. Specify filter name.")
	}

	data, err := models.GetUserProperties(ctx, start, end, filter)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, data)
}

func CheckTime(c *gin.Context) string {
	if c.Query("start") == "" && c.Query("end") == "" {
		return "Date Error: Start and End Date param missing"
	}
	return ""
}

func ParseTime(t string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func GetQueryParams(c *gin.Context, id string) string {
	param := c.Query(id)
	return param
}

func LogErrorGin(c *gin.Context, err error) {
	log.Errorf(c, "could not put into datastore: %v", err)
	x.Println(err)
	c.String(http.StatusOK, "-1")
}
