package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var avatarId = flag.String("id", "11", "头像id")
var avatarPng = flag.String("png", "icon_junshisuipian_zhangbao.png", "头像图片")

func main(){
	flag.Parse()
	url := "https://ss0-data.youzu.com/role/addavatar"
	id, err := strconv.Atoi(*avatarId)
	if err != nil {
		fmt.Println("avatarId convert err：", err)
		return
	}
	avatarPng := "https://cdn-zs-studio.uuzuonline.com/product-374/cdn1/sn/advertise/avatar/" + *avatarPng
	data := make(map[string]interface{})
	data["avatar_id"] = id
	data["avatar_url"] = avatarPng

    res,err := httpPosts(url,data)
    if err != nil {
    	fmt.Println("avatar http post error",err)
	}
	fmt.Println(res)
}

func httpPosts(url string, data interface{}) (string, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 20 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}