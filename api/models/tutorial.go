package models

import (
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Tutorial is data for us to know when to stop
type Tutorial struct {
	ID         int64     `json:"id" datastore:"-"`
	UserID     int64     `json:"userId"`
	TutorialID int64     `json:"tutorialId"`
	Date       time.Time `json:"date"`
}

func (tutorial *Tutorial) key(c context.Context) *datastore.Key {
	if tutorial.ID == 0 {
		return datastore.NewIncompleteKey(c, "Tutorial", nil)
	}
	return datastore.NewKey(c, "Tutorial", "", tutorial.ID, nil)
}

func (tutorial *Tutorial) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(tutorial)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, tutorial.key(c), tutorial)
	if err != nil {
		return err
	}

	tutorial.ID = k.IntID()
	return nil
}

// // NewTutorial inserts a new tutorial into the datastore
// func NewTutorial(c context.Context, r io.ReadCloser) (*Tutorial, error) {

// 	var tutorial Tutorial
// 	// tutorial.Timestamp = time.Now()
// 	err := json.NewDecoder(r).Decode(&tutorial)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tutorial.ID = 0

// 	err = tutorial.save(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &tutorial, nil
// }

// NewTutorial inserts a new tutorial into the datastore
func NewTutorial(c context.Context, tutorial Tutorial) (*Tutorial, error) {
	err := tutorial.save(c)
	if err != nil {
		return nil, err
	}
	return &tutorial, nil
}

// GetAllTutorials fetches all tutorial entries from datastore
func GetAllTutorials(c context.Context) ([]Tutorial, error) {
	q := datastore.NewQuery("Tutorial").Order("UserID")

	var tutorials []Tutorial
	_, err := q.GetAll(c, &tutorials)
	if err != nil {
		return nil, err
	}

	return tutorials, nil
}

// GetTutorials fetches all user property entries from datastore
func GetTutorials(c context.Context, start time.Time, end time.Time) ([]Tutorial, error) {
	q := datastore.NewQuery("Tutorial").Filter("Date >=", start).Filter("Date <", end.Add(time.Duration(time.Hour*24))).Order("Date")
	var tutorials []Tutorial
	keys, err := q.GetAll(c, &tutorials)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		tutorials[i].ID = keys[i].IntID()
	}

	if tutorials == nil {
		return []Tutorial{}, nil
	}

	return tutorials, nil
}
