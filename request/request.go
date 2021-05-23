package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get .
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

// Post .
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

	// fmt.Printf("%+v\n", request)
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
