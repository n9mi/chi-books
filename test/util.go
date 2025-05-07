package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
)

func convertReqToString(m map[string]interface{}) string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

func newRequest(method string, url string, reqBody map[string]interface{}) *http.Request {
	var req *http.Request

	if reqBody != nil {
		strReqBody := convertReqToString(reqBody)
		req = httptest.NewRequest(method, url, strings.NewReader(strReqBody))
	} else {
		req = httptest.NewRequest(method, url, nil)
	}

	req.Header.Add("Content-Type", "application/json")
	return req
}
