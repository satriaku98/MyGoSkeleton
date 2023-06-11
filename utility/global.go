package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CheckStrings(string1, string2 string) string {
	if string1 == string2 {
		return "string equals"
	}
	return "string not equals"
}
func Task() {
	currentTime := time.Now()
	fmt.Println("I am running task.", currentTime.Local().Format("2006/01/02 15:04:05"))
}
func HttpPost(requestAny []byte, urlPost string, params map[string]string) (any, error) {
	var err error
	var client = &http.Client{}
	var returnValue any

	var payload = bytes.NewBuffer(requestAny)

	request, err := http.NewRequest("POST", urlPost, payload)
	if err != nil {
		return returnValue, err
	}

	for k, v := range params {
		request.Header.Set(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return returnValue, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&returnValue)

	if err != nil {
		return returnValue, err
	}

	return returnValue, nil
}
