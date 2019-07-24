package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func JsonGet(url string, i interface{}) (err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return reqErr
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	jsonErr := json.Unmarshal(body, &i)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}
