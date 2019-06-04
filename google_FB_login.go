package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	simplejson "go-simplejson/simplejson"
)

type response_json struct {
	appid  string
	sdkuid string
}

func login(param string) (*response_json, error) {

	app_id := "951686768358766"
	app_secret := "c23c15227715e7b2d564e76746a6063e"
	grant_type := "client_credentials"
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10801")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}

	getAccessTokenUrl := "https://graph.facebook.com/oauth/access_token?client_id=" + app_id + "&client_secret=" + app_secret + "&grant_type=" + grant_type

	resp, err := client.Get(getAccessTokenUrl)
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

	//获取json的内容
	jsRes, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	tokenRes, err := jsRes.Get("access_token").String()
	if err != nil {
		return nil, err
	}

	verifyUrl := "https://graph.facebook.com/v3.2/debug_token?input_token=" + param + "&access_token=" + tokenRes

	respVerify, err := client.Get(verifyUrl)
	if err != nil {
		return nil, err
	}

	defer respVerify.Body.Close()

	bodyVerify, err := ioutil.ReadAll(respVerify.Body)
	if err != nil {
		return nil, err
	}

	//获取json的内容
	jsResVerify, err := simplejson.NewJson(bodyVerify)
	if err != nil {
		return nil, err
	}

	if jsResVerify == nil {
		return nil, err
	}

	response := response_json{}
	//获取appid
	appidRes, err := jsResVerify.Get("data").Get("app_id").String()
	if err != nil {
		return nil, err
	}
	//获取userid
	useridRes, err := jsResVerify.Get("data").Get("user_id").String()
	if err != nil {
		return nil, err
	}
	response.appid = appidRes
	response.sdkuid = useridRes

	return &response, err
}

func main() {
	var token string
	token = "EAANhjdfD9W4BADAgOlV7EMfbxYWERTBRb7EcqaEA4X6r2cUwxYo9FMWY3wh1PW1y6gxRNc5X9u2oCLVhYPDPPnE3GYsZB4jFM70VXe5vvdpPTSGNleOx8WZBu4TJaXjUHmGi5IkmVF8ZCfjkVAsKqJ0HjGbhZAAees9nwvPkwjO2MXw5Bk28WZC6vycNECDgF9zFJ0EVXXCw6ISOKJ8VO8wDlZBR7tlD4ZD"
	ret, err := login(token)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return
	}

	if ret == nil {
		fmt.Println("ret empty", err.Error())
		return
	}
}
