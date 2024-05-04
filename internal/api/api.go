package api

import (
	"fmt"
	"net/http"
)

// Api holds public api methods.
type Api interface {
	IsWebsiteStillWorking(url string) bool
}

// api holds its signature.
type api struct{}

// NewAPI starts api layer.
func NewAPI() Api {
	return &api{}
}

// IsWebsiteStillWorking checks website positive answer.
func (api) IsWebsiteStillWorking(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("fail to get resquest:", err.Error())
		return true
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
