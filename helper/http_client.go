package helper

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	http.Client
}

var freeProxy = []string{
	`80.67.195.201:81`,
	`149.248.59.104:8080`,
	`51.77.215.51:3128`,
	`208.108.120.58:8080`,
	`51.15.166.107:3128`,
	`85.10.197.9:3128`,
	`45.77.249.15:8080`,
}

func NewHttpClient() *HttpClient {
	trans := &http.Transport{
		Proxy: func(request *http.Request) (u *url.URL, err error) {

			for _, p := range freeProxy {
				u, err = url.Parse(`http://` + p)
				if err != nil {
					log.Println(err)
					continue
				}
				break
			}
			return u, err
		},
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{
		Transport: trans,
		//Timeout: 3*time.Second,
	}
	return &HttpClient{
		client,
	}
}

func (c *HttpClient) Request(method, url string, timeout time.Duration) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("new request get fail " + err.Error())
	}
	// timeout
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	req.WithContext(ctx)
	return c.Do(req)
}
