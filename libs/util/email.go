package util

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"gout/libs/setting"
)

func SendMail(params map[string]string) (bool, error) {
	conf := setting.Mail
	apiURL, ok := conf["apiURL"]
	if ok {
		delete(conf, "apiURL")
	}
	conf["subject"] = "Hello, This is an official email from BDOS Update Center"
	conf["to"] = params["email"]
	conf["html"] = params["html"]

	query := url.Values{}
	for k, v := range conf {
		query.Add(k, v)
	}

	URL := fmt.Sprintf("%s?%s", apiURL, query.Encode())
	options := map[string]interface{}{
		"headers": map[string]string{
			"Content-Type": "application/x-www-form-urlencode",
		},
	}

	r := Request{URL: URL, Options: options}
	json, err := r.Post()
	if err != nil {
		return false, err
	}
	status := json["result"].(bool)

	return status, nil
}

func RandPassword(n int) string {
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func ValidCode(n int) string {
	const letterBytes = "1234567890"
	b := make([]byte, n)
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
