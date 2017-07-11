package models

import "google.golang.org/appengine/datastore"

// DatastoreModel is the interface for datastore models
type DatastoreModel interface {
	save() error
	key() *datastore.Key
}
