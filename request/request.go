package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Get(url string) string {
	request, e := http.Get(url)
	if e != nil {
		return ""
	}
	defer request.Body.Close()
	body, e := ioutil.ReadAll(request.Body)
	if e != nil {
		return ""
	}
	return string(body)
}

func Post(u string, params map[string]string, headers ...map[string]string) string {
	v := url.Values{}
	for kk, vv := range params {
		v.Set(kk, vv)
	}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{}
	request, e := http.NewRequest("POST", u, body)
	if e != nil {
		return ""
	}
	if len(headers) > 0 {
		for kkk, vvv := range headers[0] {
			request.Header.Set(kkk, vvv)
		}
	} else {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	response, e := client.Do(request)
	if e != nil {
		return ""
	}
	defer response.Body.Close()
	data, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return ""
	}
	return string(data)
}

func PostJSON(u string, params map[string]interface{}, headers ...map[string]string) string {
	body, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	request, e := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if e != nil {
		return ""
	}

	if len(headers) > 0 {
		for kkk, vvv := range headers[0] {
			request.Header.Set(kkk, vvv)
		}
	}
	request.Header.Set("Content-Type", "application/json")
	response, e := client.Do(request)
	if e != nil {
		return ""
	}
	defer response.Body.Close()
	data, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return ""
	}
	return string(data)
}
