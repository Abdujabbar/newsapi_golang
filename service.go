package newsapi_golang

import (
	"bytes"
	"net/http"
	"net/url"
)

//NewsAPIClient struct
type NewsAPIClient struct {
	URL     string
	AuthKey string
}

//InitNewsAPIClient method
func InitNewsAPIClient(u, auth string) *NewsAPIClient {
	rurl, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	client := NewsAPIClient{URL: rurl.String(), AuthKey: auth}
	return &client
}

//GetTopHeadlines method
func (a *NewsAPIClient) GetTopHeadlines(params map[string]string) (*APIResponseArticles, error) {
	payload := url.Values{}
	for k, v := range params {
		payload.Add(k, v)
	}
	payload.Add("apiKey", a.AuthKey)

	resp, err := http.Get(a.URL + "top-headlines?" + payload.Encode())
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	body := buffer.Bytes()

	var apiResponse APIResponseArticles

	apiResponse.UnmarshalJSON(body)

	return &apiResponse, nil
}

//GetEverything method
func (a *NewsAPIClient) GetEverything(params map[string]string) (*APIResponseArticles, error) {
	payload := url.Values{}
	for k, v := range params {
		payload.Add(k, v)
	}
	payload.Add("apiKey", a.AuthKey)

	resp, err := http.Get(a.URL + "everything?" + payload.Encode())
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	body := buffer.Bytes()

	var apiResponse APIResponseArticles

	apiResponse.UnmarshalJSON(body)

	return &apiResponse, nil
}

//GetSources method
func (a *NewsAPIClient) GetSources() (*APIResponseSources, error) {
	payload := url.Values{}
	payload.Add("apiKey", a.AuthKey)
	resp, err := http.Get(a.URL + "sources?" + payload.Encode())
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	body := buffer.Bytes()

	var apiResponse APIResponseSources

	apiResponse.UnmarshalJSON(body)

	return &apiResponse, nil
}
