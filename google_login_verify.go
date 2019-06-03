package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	simplejson "go-simplejson/simplejson"
)

func login(param string) (*simplejson.Json, error) {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10801")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}

	verifyUrl := "https://oauth2.googleapis.com/tokeninfo?id_token="

	fmt.Println(verifyUrl)

	resp, err := client.Get(verifyUrl + param)
	//测试
	//resp, err := http.Get("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=monline_7_dg&wd=%20go%20dial%20tcp%20172.217.31.234%3A443%3A%20i%2Fo%20timeout&oq=dial%2520tcp%2520172.217.%2526lt%253B1.2%2526lt%253B4%253A44%2526lt%253B%253A%2520i%252Fo%2520timeout&rsv_pq=e5b61eba0005f09c&rsv_t=7b79fSPeqnzWZO2bEZyFivlvzoCNdXLWNHQ3gdvts32%2BDLKIVEb0K00%2BzcvHb3yYz7a4&rqlang=cn&rsv_enter=0&inputT=2190&rsv_sug3=6&rsv_sug2=0&rsv_sug4=3318")
	if resp == nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	/*
			if body != nil {
				fmt.Println(string(body))
		   }
	*/

	if body == nil {
		return nil, err
	}

	//获取json的内容
	jsRes, err := simplejson.NewJson(body)
	/*
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
	*/
	return jsRes, err
}

func main() {
	var token string
	token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImM3ZjUyMmQwMzIyODRkMjUyYmVlNGZkODA1NjBjZWZhMGZiNjBjMzkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyMDMwODQ5OTcwMjMtN2pybHEydWZqY2ZnMjloNHF0aGxnOGRsbG5hZmVlMWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyMDMwODQ5OTcwMjMtcnZna2NrZW1hZzE1ajJtZ2MzOGdwYm1oNmhlM25wNXMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTQ1MjI5ODY2NTQwMTYzODQzMTYiLCJlbWFpbCI6Impha3kxNjg5OTlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiLpg63lrpflsbEiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDUuZ29vZ2xldXNlcmNvbnRlbnQuY29tLy14MU5yN1p2UHBrby9BQUFBQUFBQUFBSS9BQUFBQUFBQUFKdy96N05CLUVrOHNHNC9zOTYtYy9waG90by5qcGciLCJnaXZlbl9uYW1lIjoi5a6X5bGxIiwiZmFtaWx5X25hbWUiOiLpg60iLCJsb2NhbGUiOiJ6aC1UVyIsImlhdCI6MTU1OTUzODk1NSwiZXhwIjoxNTU5NTQyNTU1fQ.ZZ5uZM7bqXxJlZuL_l9fcel-N_YBe1Oxe392znrYsnl5JRQOOjM5GGryEMilQsJf2NNZPD4Y0eXDzyeJ_tVG7j9g7n7ixqnbeseAKS1HmyoaHaFlRuf2zd4z4KfL6B9QgSCF9p9NAqINz-cky2IJ5hR85Ke_ty6cM1mKZ3atNu8-_ZbFHz5_-5goJkm_GRELfALjNbslc4tYTt3B37Ljc9KCx0ovvAuWy7qroOoO1SnTzjyqO1dlL8RNdmMRiNUTI_b7HJEBVnZSG6mh92-V-n0Z_XT38ZPLIS5D9TJCw-lQRlAU0w82jyJN-xameWCuADGSx09Ki0J5L-uBVhXeoQ"
	jsRes, err := login(token)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return
	}

	if jsRes == nil {
		fmt.Println("jsRes empty", err.Error())
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
