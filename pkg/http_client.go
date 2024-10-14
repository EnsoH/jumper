package pkg

import "net/http"

func New() *http.Client {
	return &http.Client{}
}
