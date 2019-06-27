// Package google provides a function to do Google searches using the Google Web
// Search API. See https://developers.google.com/web-search/docs/
//
// This package is an example to accompany https://blog.golang.org/context.
// It is not intended for use by others.
//
// Google has since disabled its search API,
// and so this package is no longer useful.
package google

import (
	"context"
	"encoding/json"
	"fmt"
	"go-example/context/searchApi/userip"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Results is an ordered list of search results.
type Results []Result

// A Result contains the title and URL of a search result.
type Result struct {
	Title, URL string
}

type Querys map[string]string

// Search sends query to Google search and returns the results.
func Search(ctx context.Context, qs Querys) (Results, error) {
	// Prepare the Google Search API request.  More document see https://developers.google.com/custom-search/v1/cse/list
	key := `AIzaSyBwVAtrthVBM-53peyJmSc5PKqw1Q9Z_2Y`
	cx := `008943617856372985247:bkqnk_r8-dk`
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s", key, cx)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	// Add Query Strings
	for qKey, qVal := range qs {
		q.Set(qKey, qVal)
	}

	// If ctx is carrying the user IP address, forward it to the server.
	// Google APIs use the user IP to distinguish server-initiated requests
	// from end-user requests.
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	// Issue the HTTP request and handle the response. The httpDo function
	// cancels the request if ctx.Done is closed.
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result.
		// https://developers.google.com/web-search/docs/#fonje
		var data struct {
			SearchInfo struct {
				SearchTime   float64 `json:"searchTime"`
				TotalResults string  `json:"totalResults"`
			} `json:"searchInformation"`
			Results []struct {
				Title   string `json:"title"`
				Link    string `json:"link"`
				Snippet string `json:"snippet"`
			} `json:"items"`
		}

		bs, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(bs))
		if err := json.Unmarshal(bs, &data); err != nil {
			return err
		}
		for _, res := range data.Results {
			results = append(results, Result{Title: res.Title, URL: res.Link})
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here.
	return results, err
}

// httpDo issues the HTTP request and calls f with the response. If ctx.Done is
// closed while the request or f is running, httpDo cancels the request, waits
// for f to exit, and returns ctx.Err. Otherwise, httpDo returns f's error.
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)

	// Get HTTP Client
	client, err := getHttpClient()
	if err != nil {
		return err
	}

	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}

// If google.com prevent, you can set a sock5 proxy
func getHttpClient() (*http.Client, error) {
	// whether need http proxy
	_, err := net.DialTimeout("tcp", "google.com", 1*time.Second)
	if err == nil {
		return http.DefaultClient, nil
	}

	// http transport dial content set
	httpTransport := http.Transport{}
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:10443", nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		return dialer.Dial(network, addr)
	}

	// http proxy client
	return &http.Client{Transport: &httpTransport}, nil
}
