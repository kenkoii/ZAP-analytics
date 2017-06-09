package handlers

import (
	"encoding/json"
	x "log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/kenkoii/Analytics/api/models"
)

// PostEntriesEndpoint handles POST requests on analytics endpoint
func PostEntriesEndpoint(w http.ResponseWriter, r *http.Request) {
	var count = 0
	ctx := appengine.NewContext(r)
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
