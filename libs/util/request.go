package util

import (
	"encoding/json"
	"net/http"

	"github.com/mozillazg/request"
)

type Request struct {
	URL     string
	Options map[string]interface{}
}

func (r *Request) Get() (map[string]interface{}, error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = r.Options["headers"].(map[string]string)
	resp, err := req.Get(r.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	j, err := resp.Content()

	var result map[string]interface{}
	if err := json.Unmarshal(j, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Request) Post() (map[string]interface{}, error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = r.Options["headers"].(map[string]string)
	if body, ok := r.Options["body"]; ok {
		req.Json = body
	}
	resp, err := req.Post(r.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	j, err := resp.Content()

	var result map[string]interface{}
	if err := json.Unmarshal(j, &result); err != nil {
		return nil, err
	}

	return result, nil
}
