package models

import (
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// UserPurchase is information regarding each purchase made
type UserPurchase struct {
	ID       int64     `json:"id" datastore:"-"`
	UserID   int64     `json:"userId"`
	ItemID   int64     `json:"itemId"`
	Date     time.Time `json:"date"`
	Price    float64   `json:"price"`
	IsFirst  bool      `json:"isFirst"`
	Currency string    `json:"currency"`
}

func (userPurchase *UserPurchase) key(c context.Context) *datastore.Key {
	if userPurchase.ID == 0 {
		return datastore.NewIncompleteKey(c, "UserPurchase", nil)
	}
	return datastore.NewKey(c, "UserPurchase", "", userPurchase.ID, nil)
}

func (userPurchase *UserPurchase) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(userPurchase)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, userPurchase.key(c), userPurchase)
	if err != nil {
		return err
	}

	userPurchase.ID = k.IntID()
	return nil
}

// NewUserPurchase inserts a new entry into the datastore
// func NewUserPurchase(c context.Context, r io.ReadCloser) (*UserPurchase, error) {
// 	var userPurchase UserPurchase
// 	// userPurchase.Timestamp = time.Now()
// 	err := json.NewDecoder(r).Decode(&userPurchase)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userPurchase.ID = 0

// 	err = userPurchase.save(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userPurchase, nil
// }

func NewUserPurchase(c context.Context, userPurchase UserPurchase) (*UserPurchase, error) {
	err := userPurchase.save(c)
	if err != nil {
		return nil, err
	}
	return &userPurchase, nil
}

// GetAllUserPurchases fetches all user property entries from datastore
func GetAllUserPurchases(c context.Context) ([]UserPurchase, error) {
	q := datastore.NewQuery("UserPurchase").Order("UserID")

	var userPurchases []UserPurchase
	_, err := q.GetAll(c, &userPurchases)
	if err != nil {
		return nil, err
	}

	return userPurchases, nil
}

// GetUserPurchases fetches all user property entries from datastore
func GetUserPurchases(c context.Context, start time.Time, end time.Time) ([]UserPurchase, error) {
	q := datastore.NewQuery("UserPurchase").Filter("Date >=", start).Filter("Date <", end.Add(time.Duration(time.Hour*24))).Order("Date")
	var userPurchases []UserPurchase
	keys, err := q.GetAll(c, &userPurchases)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		userPurchases[i].ID = keys[i].IntID()
	}

	if userPurchases == nil {
		return []UserPurchase{}, nil
	}

	return userPurchases, nil
}
