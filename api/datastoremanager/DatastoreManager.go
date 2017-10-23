package datastoremanager

import (
	"github.com/gin-gonic/gin"
	"github.com/kenkoii/Analytics/api/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func GetUserProperties(g *gin.Context) ([]models.UserProperty, error) {
	c := appengine.NewContext(g.Request)
	query := datastore.NewQuery("UserProperty").Order("UserID")
	var userProperties []models.UserProperty
	_, err := query.GetAll(c, &userProperties)
	if err != nil {
		return nil, err
	}
	return userProperties, nil
}
