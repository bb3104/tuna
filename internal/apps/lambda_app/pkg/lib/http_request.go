package lib

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}
