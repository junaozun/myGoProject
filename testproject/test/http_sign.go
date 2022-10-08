package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var addr = "http://127.0.0.1:14001?mod=vms&r=gameApi/getMaintain"
var key = "Tt7TOZwUqYOI6lBuWyjjaB0zwbdnWt72"

type VmsInfo struct {
}

func main() {
	resp := &VmsInfo{}

	// post
	reqMap := map[string]interface{}{
		"opgameId": 1000,
	}
	err := postWithSign(addr, reqMap, resp)
	if err != nil {
		panic(err)
	}

	// postform
	// dataUrlVal := url.Values{}
	// dataUrlVal.Add("sid", "10040111122")
	// dataUrlVal.Add("account", "nihao")
	// dataUrlVal.Add("gate", "11223")
	// err := postFormWithSign(addr, dataUrlVal, resp)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(resp)
}

func postWithSign(addr string, data map[string]interface{}, out interface{}) error {
	dataUrlVal := url.Values{}
	for key, val := range data {
		mm, err := ParseString(val)
		if err != nil {
			return err
		}
		dataUrlVal.Add(key, mm)
	}
	sign, err := calculateSignSHA(dataUrlVal, key)
	if err != nil {
		return err
	}
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sign", sign)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func postFormWithSign(addr string, data url.Values, out interface{}) error {
	sign, err := calculateSignSHA(data, key)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", addr, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	// 当使用url.Values传入request时，必须使用"application/x-www-form-urlencoded".服务器解析时，c.PostForm("key")中才会有值
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("sign", sign)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func calculateSignSHA(values url.Values, key string) (string, error) {
	decodeStr, err := url.QueryUnescape(values.Encode())
	if nil != err {
		return "", err
	}
	h := sha1.New()
	h.Write([]byte(decodeStr + key))
	bs := h.Sum(nil)
	return strings.ToLower(fmt.Sprintf("%x", bs)), nil
}

func ParseString(t interface{}) (string, error) {
	switch value := t.(type) {
	case string:
		return value, nil
	case float64:
		return strconv.FormatFloat(value, 'f', 0, 64), nil
	case int:
		return strconv.Itoa(value), nil
	default:
		return "", fmt.Errorf("unknown type")
	}
}

// VerifySign 兼容postForm和post两种方式的数据提交验证
func VerifySign(c *gin.Context) error {
	sm := c.Query("mod")
	if sm != "vms" {
		return errors.New("不是vms接口")
	}
	c.Request.ParseForm()
	valueMap := make(map[string][]string)
	// 先判断下是否postForm数据
	postForm := c.Request.PostForm
	// postForm数据验证
	if len(postForm) != 0 {
		valueMap = postForm
	} else { // post数据验证
		b, _ := c.GetRawData()
		if len(b) != 0 {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			tempMap := make(map[string]interface{})
			err := json.Unmarshal(b, &tempMap)
			if err != nil {
				return errors.New("参数解析错误")
			}
			for k, v := range tempMap {
				value, err := ParseString(v)
				if err != nil {
					return errors.New("参数解析错误")
				}
				valueMap[k] = []string{value}
			}
		}
	}
	sign, err := calculateSignSHA(valueMap, key)
	if err != nil {
		return errors.New("签名计算错误")
	}

	reqSign := c.Request.Header.Get("sign")
	if sign != reqSign {
		return errors.New("签名验证不匹配")
	}
	return nil
}
