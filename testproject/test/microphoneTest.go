package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 添加页面
func addhandler(w http.ResponseWriter, r *http.Request) {

	// 如果是post请求，注意先后顺序
	if r.Method == http.MethodPost {
		// 解析表单 并且校验
		r.ParseForm()
		tokenId := r.FormValue("tokenId")
		quality := r.FormValue("quality")
		tokenIdN, err := strconv.ParseInt(tokenId, 10, 64)
		if err != nil {
			fmt.Fprint(w, "strconv.ParseInt err")
			return
		}
		tt := AddReq{
			TokenId: tokenIdN,
			Quality: quality,
		}
		b, err := json.Marshal(tt)
		if err != nil {
			panic(err)
		}
		header := map[string]string{
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjM3NjUzNDUsInVzZXJOYW1lIjoi55So5oi3Mzc2NTM0NSIsImV4cCI6MTY2MTQ5MzY2MCwiaXNzIjoieHgifQ.lJJKJxUP0NZkPpNuEKs1Pz3ZTqlS3kfEVdZXPPQDg6s",
		}
		res := httpPostJsonString("http://47.98.211.160/lobby/recharge_microphone_Test", header, string(b))

		ooo := &AddResp{}
		err = json.Unmarshal(res, ooo)
		if err != nil {
			panic(err)
		}
		t, err := template.ParseFiles("add.html")
		if err != nil {
			log.Fatal(err)
		}

		data := map[string]interface{}{
			"message":   ooo.ErrMsg,
			"maxPower":  ooo.Data.MaxPower,
			"maxIncome": ooo.Data.MaxIncome,
			"lockStage": ooo.Data.LockStage,
			"tokenIds":  ooo.Data.TokenIds,
			"income":    ooo.Data.Income,
		}
		// 渲染模板
		t.Execute(w, data)
		// io.WriteString(w, string(res))
	} else {
		// 解析模板
		t, err := template.ParseFiles("add.html")
		if err != nil {
			log.Fatal(err)
		}
		// 渲染模板
		t.Execute(w, nil)
	}

}

func deletehandler(w http.ResponseWriter, r *http.Request) {
	// 如果是post请求，注意先后顺序
	if r.Method == http.MethodPost {
		// 解析表单 并且校验
		r.ParseForm()
		tokenId := r.FormValue("tokenId")
		tokenIdN, err := strconv.ParseInt(tokenId, 10, 64)
		if err != nil {
			fmt.Fprint(w, "strconv.ParseInt err")
			return
		}
		tt := DelReq{
			TokenId: tokenIdN,
		}
		b, err := json.Marshal(tt)
		if err != nil {
			panic(err)
		}
		header := map[string]string{
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjM3NjUzNDUsInVzZXJOYW1lIjoi55So5oi3Mzc2NTM0NSIsImV4cCI6MTY2MTQ5MzY2MCwiaXNzIjoieHgifQ.lJJKJxUP0NZkPpNuEKs1Pz3ZTqlS3kfEVdZXPPQDg6s",
		}
		res := httpPostJsonString("http://47.98.211.160/lobby/delete_microphone_Test", header, string(b))

		ooo := &AddResp{}
		err = json.Unmarshal(res, ooo)
		if err != nil {
			panic(err)
		}
		t, err := template.ParseFiles("delete.html")
		if err != nil {
			log.Fatal(err)
		}

		data := map[string]interface{}{
			"message":   ooo.ErrMsg,
			"maxPower":  ooo.Data.MaxPower,
			"maxIncome": ooo.Data.MaxIncome,
			"lockStage": ooo.Data.LockStage,
			"tokenIds":  ooo.Data.TokenIds,
			"income":    ooo.Data.Income,
		}
		// 渲染模板
		t.Execute(w, data)
		// io.WriteString(w, string(res))
	} else {
		// 解析模板
		t, err := template.ParseFiles("delete.html")
		if err != nil {
			log.Fatal(err)
		}
		// 渲染模板
		t.Execute(w, nil)
	}
}

func addIncome(w http.ResponseWriter, r *http.Request) {
	// 如果是post请求，注意先后顺序
	if r.Method == http.MethodPost {
		// 解析表单 并且校验
		r.ParseForm()
		income := r.FormValue("income")
		incomeN, err := strconv.ParseInt(income, 10, 64)
		if err != nil {
			fmt.Fprint(w, "strconv.ParseInt err")
			return
		}
		tt := IncomeReq{
			Income: incomeN,
		}
		b, err := json.Marshal(tt)
		if err != nil {
			panic(err)
		}
		header := map[string]string{
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjM3NjUzNDUsInVzZXJOYW1lIjoi55So5oi3Mzc2NTM0NSIsImV4cCI6MTY2MTQ5MzY2MCwiaXNzIjoieHgifQ.lJJKJxUP0NZkPpNuEKs1Pz3ZTqlS3kfEVdZXPPQDg6s",
		}
		res := httpPostJsonString("http://47.98.211.160/lobby/add_CurIncomeDay_Test", header, string(b))

		ooo := &IncomeResp{}
		err = json.Unmarshal(res, ooo)
		if err != nil {
			panic(err)
		}
		var code string
		if ooo.ErrCode != 0 {
			code = "发起请求失败"
		} else {
			code = "添加收益成功"
		}
		t, err := template.ParseFiles("income.html")
		if err != nil {
			log.Fatal(err)
		}

		data := map[string]interface{}{
			"message": ooo.ErrMsg,
			"code":    code,
		}
		// 渲染模板
		t.Execute(w, data)

	} else {
		// 解析模板
		t, err := template.ParseFiles("income.html")
		if err != nil {
			log.Fatal(err)
		}
		// 渲染模板
		t.Execute(w, nil)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", addhandler)
	mux.HandleFunc("/delete", deletehandler)
	mux.HandleFunc("/addIncome", addIncome)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func httpPostJsonString(strURL string, mapHeader map[string]string, strJsonData string) []byte {
	client := &http.Client{}

	body := strings.NewReader(strJsonData)

	req, err := http.NewRequest("POST", strURL, body)
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.FormatInt(body.Size(), 10))
	// req.Header.Set("Connection", "keep-alive")
	if mapHeader != nil {
		for k, v := range mapHeader {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	return respBody
}

type AddReq struct {
	TokenId int64  `json:"tokenId"`
	Quality string `json:"quality"`
}

type DelReq struct {
	TokenId int64 `json:"tokenId"`
}

type IncomeReq struct {
	Income int64 `json:"income"`
}

type AddResp struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    struct {
		MaxPower  int `json:"maxPower"`
		MaxIncome int `json:"maxIncome"`
		LockStage []struct {
			Progress   int `json:"progress"`
			Status     int `json:"status"`
			UnLockUid  int `json:"unLockUid"`
			UnLockTime int `json:"unLockTime"`
		} `json:"lockStage"`
		TokenIds []int64 `json:"tokenIds"`
		Income   int64   `json:"income"`
	} `json:"data"`
}

type IncomeResp struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}
