package api

import (
	"fmt"
	"net/http"
)

func IsWebsiteStillWorking(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("fail to get resquest:", err.Error())
		return true
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
