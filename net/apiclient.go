/*
ApiClient

Licensed under MIT license
(c) 2017 Reeleezee BV
*/
package net

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// --------------------------------------------------------------------
// ApiClient, API Client struct
// --------------------------------------------------------------------
type ApiClient struct {
	client        http.Client
	authorization string
	uri           string
	IsJSON        bool
}

func createRequest(method, urlStr, authorization string, body []byte) *http.Request {
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Basic "+authorization)
	if method == "PUT" || method == "POST" {
		req.Header.Add("Prefer", "return=representation")
	}
	return req
}

func (ac *ApiClient) handleRequest(method string, route string, data []byte) (int, []byte) {
	req := createRequest(method, ac.uri+route, ac.authorization, data)
	resp, err := ac.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	ac.IsJSON = strings.Contains(resp.Header.Get("Content-Type"), "application/json")
	body, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}

func (ac *ApiClient) Init(uri string, userName string, password string) {
	ac.uri = uri
	ac.authorization = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userName, password)))
	ac.client.Timeout = 90 * time.Second
}

func (ac *ApiClient) Get(route string) (int, []byte) {
	return ac.handleRequest("GET", route, nil)
}

func (ac *ApiClient) Put(route string, data []byte) (int, []byte) {
	return ac.handleRequest("PUT", route, data)
}
