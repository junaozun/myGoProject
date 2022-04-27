package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	contentType = "application/x-www-form-urlencoded"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func PostHttp(url string, data interface{}) (string, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 20 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}

func httpPost(addr string, data map[string]string) {
	dataUrlVal := url.Values{}
	for key, val := range data {
		dataUrlVal.Add(key, val)
	}
	resp, err := http.Post(addr, contentType, strings.NewReader(dataUrlVal.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

// "application/x-www-form-urlencoded"
func PostForm(url string, value url.Values, out interface{}) error {
	client := http.Client{
		Timeout: 3 * time.Second, // todo 把不同接口的限制超时时间都定义在一块
	}

	resp, err := client.PostForm(url, value)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respData, out)
}

// test
func main() {
	//post要提交的数据
	dataUrlVal := url.Values{}
	dataUrlVal.Add("sid", "10040111122")
	dataUrlVal.Add("account", "nihao")
	dataUrlVal.Add("gate", "11223")
	PostForm("http://127.0.0.1:10008/check_login", dataUrlVal, nil)
}
