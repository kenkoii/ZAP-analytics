package models

import (
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Gacha is data for us to determine if the gacha is balance
type Gacha struct {
	ID         int64     `json:"id" datastore:"-"`
	UserID     int64     `json:"userId"`
	Date       time.Time `json:"date"`
	GachaID    int64     `json:"gachaId"`
	IsTen      bool      `json:"isTen"`
	IsFree     bool      `json:"isFree"`
	IsPurchase bool      `json:"isPurchase"`
	PicupID    int64     `json:"picupId"`
}

func (gacha *Gacha) key(c context.Context) *datastore.Key {
	if gacha.ID == 0 {
		return datastore.NewIncompleteKey(c, "Gacha", nil)
	}
	return datastore.NewKey(c, "Gacha", "", gacha.ID, nil)
}

func (gacha *Gacha) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(gacha)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, gacha.key(c), gacha)
	if err != nil {
		return err
	}

	gacha.ID = k.IntID()
	return nil
}

// NewStage inserts a new gacha into the datastore
func NewGacha(c context.Context, gacha Gacha) (*Gacha, error) {
	err := gacha.save(c)
	if err != nil {
		return nil, err
	}
	return &gacha, nil
}

// GetAllGachas fetches all gacha entries from datastore
func GetAllGachas(c context.Context) ([]Gacha, error) {
	q := datastore.NewQuery("Gacha").Order("UserID")

	var gachas []Gacha
	_, err := q.GetAll(c, &gachas)
	if err != nil {
		return nil, err
	}

	return gachas, nil
}

// GetGachas fetches all user property entries from datastore
func GetGachas(c context.Context, start time.Time, end time.Time) ([]Gacha, error) {
	q := datastore.NewQuery("Gacha").Filter("Date >=", start).Filter("Date <", end.Add(time.Duration(time.Hour*24))).Order("Date")
	var gachas []Gacha
	keys, err := q.GetAll(c, &gachas)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		gachas[i].ID = keys[i].IntID()
	}

	if gachas == nil {
		return []Gacha{}, nil
	}

	return gachas, nil
}
