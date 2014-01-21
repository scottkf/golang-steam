package api

import "net/http"
import "io/ioutil"

func ApiRequest(url string) ([]byte, int) {
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	return body, res.StatusCode
}
