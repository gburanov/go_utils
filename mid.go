package goutils

import (
	"fmt"
	"net/http"
)

type cntFunc func(w http.ResponseWriter, r *http.Request) error

// HTTPError is error with error code
type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("Code %d, %s", e.Code, e.Message)
}

func HandleErrors(mid cntFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mid(w,r)
		if err != nil {
			if ae, ok := err.(*HTTPError); ok {
				http.Error(w, ae.Error(), ae.Code)
			} else {
				http.Error(w, err.Error(), 500)
			}
			return
		}
		fmt.Fprintln(w, "OK")
	}
}
