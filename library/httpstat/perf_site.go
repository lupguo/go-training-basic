package main

import (
	"fmt"
	"github.com/tcnksm/go-httpstat"
	"go-example/helper"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

// 单页
type PerfUrl struct {
	title string
	url   string
}

// 性能测试清单
type PerfSystem map[string][]PerfUrl

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	systems := PerfSystem{
		"GB_WWW": []PerfUrl{
			{"首页", "https://www.gearbest.com/"},
			{"FlashSale", "https://www.gearbest.com/flash-sale.html?top-goods=438932701-1433363"},
			{"搜索", "https://www.gearbest.com/xiaomi-_gear/"},
			{"搜索品牌", "https://www.gearbest.com/xiaomi-_gear/?brand=luanke"},
			{"搜索分类", "https://www.gearbest.com/xiaomi-_gear/c_12993/"},
			{"新品", "https://www.gearbest.com/new-products/"},
			{"Cool", "https://www.gearbest.com/explore/"},
			{"商详", "https://www.gearbest.com/makeup-brushes-tools/pp_1116016.html?wid=1433363"},
		},
		"GB_WAP": []PerfUrl{
			{"首页", "https://m.gearbest.com"},
			{"搜索列表", "https://m.gearbest.com/Xiaomi-Mi-9T-_gear/?page_size=36&odr=hot"},
			{"所有分类", "https://m.gearbest.com/all-categories.html"},
			{"指定分类", "https://m.gearbest.com/cookware-c_11628/"},
			{"分类筛选", "https://m.gearbest.com/bakeware-c_11626/?on_sale=1&price_min=1.0000&price_max=10.0000"},
			{"活动页", "https://m.gearbest.com/activity-mid-year-sale.html"},
			{"商详", "https://m.gearbest.com/cake-molds/pp_921479.html?wid=1433363"},
		},
	}

	// Create a httpstat powered context
	result := new(httpstat.Result)

	// Http request
	tw := tabwriter.NewWriter(os.Stdout, 16, 4, 4, ' ', 0)
	show := helper.TwShow(tw)
	defer tw.Flush()
	show("DNSLookup", "TCPConnection", "TLSHandshake", "ServerProcessing", "ContentTransfer", "Url", "Title", )
	reqClient := http.DefaultClient
	for _, u := range systems["GB_WWW"] {
		if err := TraceHttpRequest(reqClient, u.url, result); err != nil {
			log.Printf("%s, %s, %s\n", u.title, u.url, err)
		}
		show(
			fmt.Sprintf("%d ms", int(result.DNSLookup/time.Millisecond)),
			fmt.Sprintf("%d ms", int(result.TCPConnection/time.Millisecond)),
			fmt.Sprintf("%d ms", int(result.TLSHandshake/time.Millisecond)),
			fmt.Sprintf("%d ms", int(result.ServerProcessing/time.Millisecond)),
			fmt.Sprintf("%d ms", int(result.Pretransfer/time.Millisecond)),
			u.title,
			u.url,
		)
	}
}

func TraceHttpRequest(client *http.Client, rawUrl string, result *httpstat.Result) (err error) {
	// no cache key
	u, _ := url.Parse(rawUrl)
	uv := u.Query()
	uv.Add("no-cache", strconv.Itoa(time.Now().Nanosecond()))
	u.RawQuery = uv.Encode()

	// build request
	newUrl := u.String()
	req, err := http.NewRequest("GET", newUrl, nil);
	if err != nil {
		return err
	}
	// trace context
	ctx := httpstat.WithHTTPStat(req.Context(), result)
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// copy resp.body to /dev/null
	if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
		return
	}
	_ = resp.Body.Close()
	result.End(time.Now())
	return
}
