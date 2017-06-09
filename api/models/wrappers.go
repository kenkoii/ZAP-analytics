package models

import "net/http"
import "encoding/json"

// AnalyticsRequest is the expected structure of request body
type AnalyticsRequest struct {
	ClassName string          `json:"className"`
	Params    json.RawMessage `json:"params"`
}

// ErrorHandler is a wrapper for error handling structure of request body
type ErrorHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
