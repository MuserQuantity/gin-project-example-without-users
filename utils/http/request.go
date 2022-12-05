package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(url string, params *url.Values) ([]byte, error) {
	var paramsStr string
	if params != nil {
		paramsStr = "?" + params.Encode()
	}
	resp, err := http.Get(url + paramsStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func PostJson(url string, data interface{}) (response []byte, err error) {
	response, err = Request(url, "POST", "json", data)
	return
}

func Request(url, reqType, dataType string, data interface{}) ([]byte, error) {
	cl := &http.Client{}
	var params []byte
	var byteParam *strings.Reader
	var req *http.Request
	var err error
	// request body
	params, err = json.Marshal(data) // transform to json
	if err != nil {
		return nil, err
	}
	byteParam = strings.NewReader(string(params))
	req, err = http.NewRequest(reqType, url, byteParam)
	if err != nil {
		return nil, err
	}
	// content-type
	if dataType == "json" {
		dt := "application/json;charset=utf-8"
		req.Header.Add("Content-Type", dt)
	} else if dataType == "form" {
		dt := "application/x-www-form-urlencoded"
		req.Header.Add("Content-Type", dt)
	}
	// do request
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// get response
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%s", response)
	}
	return response, nil
}
