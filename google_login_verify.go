package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	simplejson "go-simplejson/simplejson"
)

func Indexhandler(param string) {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10801")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}

	var verifyUrl string = "https://oauth2.googleapis.com/tokeninfo?id_token="
	fmt.Println(verifyUrl)
	resp, err := client.Get(verifyUrl + param)
	//测试
	//resp, err := http.Get("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=monline_7_dg&wd=%20go%20dial%20tcp%20172.217.31.234%3A443%3A%20i%2Fo%20timeout&oq=dial%2520tcp%2520172.217.%2526lt%253B1.2%2526lt%253B4%253A44%2526lt%253B%253A%2520i%252Fo%2520timeout&rsv_pq=e5b61eba0005f09c&rsv_t=7b79fSPeqnzWZO2bEZyFivlvzoCNdXLWNHQ3gdvts32%2BDLKIVEb0K00%2BzcvHb3yYz7a4&rqlang=cn&rsv_enter=0&inputT=2190&rsv_sug3=6&rsv_sug2=0&rsv_sug4=3318")
	if resp == nil {
		fmt.Println("Fatal error", err.Error())
		return
	}

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	/*
			if body != nil {
				fmt.Println(string(body))
		   }
	*/

	if body == nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	//获取json的内容
	jsRes, err := simplejson.NewJson(body)
	if err != nil {
		fmt.Println("simplejson error", err.Error())
		return
	}

	if jsRes == nil {
		fmt.Println("jsRes error", err.Error())
		return
	}

	//根据json的格式，选择使用array或者map存储数据
	var data = make(map[string]interface{})
	data, err = jsRes.Map()
	if err != nil {
		fmt.Println("storage data error", err.Error())
		return
	}

	if data == nil {
		fmt.Println("storage data error", err.Error())
		return
	}

	fmt.Println(data["aud"])
}

func main() {
	var token string = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImM3ZjUyMmQwMzIyODRkMjUyYmVlNGZkODA1NjBjZWZhMGZiNjBjMzkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyMDMwODQ5OTcwMjMtN2pybHEydWZqY2ZnMjloNHF0aGxnOGRsbG5hZmVlMWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyMDMwODQ5OTcwMjMtcnZna2NrZW1hZzE1ajJtZ2MzOGdwYm1oNmhlM25wNXMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTA1OTUyMzY0NTk2NTQ2NjkyMjQiLCJlbWFpbCI6ImEwMjAzODM3NDVAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiLlkLPmma_mpq4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDYuZ29vZ2xldXNlcmNvbnRlbnQuY29tLy15cV90M0VJS251SS9BQUFBQUFBQUFBSS9BQUFBQUFBQUFBQS9BQ0hpM3JjNEh1UjZncjRtd3BlZFhsd0FRRDRLcFpJbWh3L3M5Ni1jL3Bob3RvLmpwZyIsImdpdmVuX25hbWUiOiLmma_mpq4iLCJmYW1pbHlfbmFtZSI6IuWQsyIsImxvY2FsZSI6InpoLVRXIiwiaWF0IjoxNTU5NDgxNjgyLCJleHAiOjE1NTk0ODUyODJ9.fgZRuUo200kA9spnH4lrLb59v0i2EdJ048Lr1X4IXnEbRAMGo-D762RJM1j48I9SumrlbdXJVckzX8phRTkG0K3Js8NlrHegCwuttJNTPQDGnoC7rd1baGnOTmBzCXBUP4S1K6a1_2rnjCriGZLhs0del57mDMzq7S6a_qSDDn4ZhQkrfFsdXT72k5cI4Hvg4AFEv49kDxJqvsWVuA9rP6YdjbGTlYqDAuxZvd4Urr_sxlk1WEuwOd7W6xxNNemgxTZK7DMRKziPNTRy8DeQFHCFAmkmuz9Uy32KCgC2eJkDOi9kT77zCQFstD-jZC5CaEkMH3OItevp9eu8TvOqxQ"
	Indexhandler(token)
}
