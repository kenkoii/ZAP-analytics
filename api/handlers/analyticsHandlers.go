package handlers

import (
	"encoding/json"
	x "log"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/kenkoii/Analytics/api/models"
)

// PostEntriesEndpoint handles POST requests on analytics endpoint
func PostEntriesEndpoint(w http.ResponseWriter, r *http.Request) {
	var count = 0
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)
	var analyticsRequests []models.AnalyticsRequest
	err := json.NewDecoder(r.Body).Decode(&analyticsRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(analyticsRequests); i++ {
		switch analyticsRequests[i].ClassName {
		case "UserProperty":
			{
				var userProperty models.UserProperty
				if err := json.Unmarshal(analyticsRequests[i].Params, &userProperty); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewUserProperty(ctx, userProperty)
				if err != nil {
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "UserPurchase":
			{
				var userPurchase models.UserPurchase
				if err := json.Unmarshal(analyticsRequests[i].Params, &userPurchase); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewUserPurchase(ctx, userPurchase)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "UserDailyProperty":
			{
				var userDailyProperty models.UserDailyProperty
				if err := json.Unmarshal(analyticsRequests[i].Params, &userDailyProperty); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewUserDailyProperty(ctx, userDailyProperty)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "Tutorial":
			{
				var tutorial models.Tutorial
				if err := json.Unmarshal(analyticsRequests[i].Params, &tutorial); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewTutorial(ctx, tutorial)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "Stage":
			{
				var stage models.Stage
				if err := json.Unmarshal(analyticsRequests[i].Params, &stage); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewStage(ctx, stage)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "Quest":
			{
				var quest models.Quest
				if err := json.Unmarshal(analyticsRequests[i].Params, &quest); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewQuest(ctx, quest)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "Event":
			{
				var event models.Event
				if err := json.Unmarshal(analyticsRequests[i].Params, &event); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewEvent(ctx, event)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		case "Gacha":
			{
				var gacha models.Gacha
				if err := json.Unmarshal(analyticsRequests[i].Params, &gacha); err != nil {
					LogError(ctx, err, w)
					return
				}

				_, err := models.NewGacha(ctx, gacha)
				if err != nil {
					// http.Error(w, err.Error(), http.StatusInternalServerError)
					LogError(ctx, err, w)
					return
				}
				count++
			}
		}
	}
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(ctx, "could not put into datastore: %v", err)
		x.Println(err)
		json.NewEncoder(w).Encode(-1)
		return
	}
	json.NewEncoder(w).Encode(1)
}

func LogError(ctx context.Context, err error, w http.ResponseWriter) {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Errorf(ctx, "could not put into datastore: %v", err)
	x.Println(err)
	json.NewEncoder(w).Encode(1)
}

// GetUserPurchaseEndpoint handles POST requests on analytics endpoint
func GetUserPurchaseEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)
	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	results, err := models.GetUserPurchases(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

func GetUserPropertiesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)
	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	filter := r.FormValue("filter")
	if filter == "" {
		http.Error(w, "Request Error: Filter param missing. Specify filter name.", http.StatusInternalServerError)
		return
	}

	results, err := models.GetUserProperties(ctx, start, end, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// GetUserDailyPropertiesEndpoint handles POST requests on analytics endpoint
func GetUserDailyPropertiesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)
	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	results, err := models.GetUserDailyProperties(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// GetUserDailyPropertiesEndpoint handles POST requests on analytics endpoint
func GetTutorialsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)
	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	results, err := models.GetTutorials(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// GetStagesEndpoint handles POST requests on analytics endpoint
func GetStagesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)

	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	results, err := models.GetStages(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

func GetGachasEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(appengine.NewContext(r), time.Second*30)

	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	results, err := models.GetGachas(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

/*


func GetDataByDateEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.FormValue("table") == "" {
		http.Error(w, "Request Error: Table param missing. Specify table name.", http.StatusInternalServerError)
		return
	}

	if r.FormValue("start") == "" && r.FormValue("end") == "" {
		http.Error(w, "Date Error: Start and End Date param missing", http.StatusInternalServerError)
		return
	}

	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "Start Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "End Date Error: Please follow format YYYY-MM-DD", http.StatusInternalServerError)
		return
	}

	switch r.FormValue("table") == "" {
		case "UserProperty":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "UserPurchase":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "UserDailyProperty":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "Tutorial":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "Stage":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "Quest":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		case "Event":
			{
				results, err := models.GetUserPurchases(ctx, start, end)
			}
		}



	results, err := models.GetUserPurchases(ctx, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

*/
