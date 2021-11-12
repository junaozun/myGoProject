package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	ContentType = "application/x-www-form-urlencoded"
)

var (
	wg       sync.WaitGroup
	client   *http.Client
	accounts []string
	count    int
)

func main() {
	fmt.Println("账号创建中......")
	client = &http.Client{Timeout: 120 * time.Second}
	// 1.设置集群时间内，报名期
	//code := WonPost("http://localhost:8080/request?account=sxf&rpc=20610", `{"sn":0,"guide":0,"trigger":"","str":"setClusterTime=2021-08-24 16:01:00"}`)
	//if code != 1 {
	//	panic("setClusterTime error" + strconv.Itoa(code))
	//}
	//
	//// 创建主账号
	//wg.Add(5)
	//go GenWonPlayerData("sxf", "神龙教1", true, []int32{20, 28})
	//go GenWonPlayerData("sxf1", "神龙教2", true, []int32{24, 28})
	//go GenWonPlayerData("sxf2", "神龙教3", true, []int32{28, 28})
	//go GenWonPlayerData("sxf3", "神龙教4", true, []int32{17, 28})
	//go GenWonPlayerData("sxf4", "神龙教5", true, []int32{7, 28})
	//wg.Wait()
	code := WonPost("http://localhost:8080/request?account=sxf&rpc=40035", `{"sn":0,"guide":0,"trigger":"","mapId":1,
	   "wonFtCoor":[
	   {
	     "ft": 44,
	     "coor":[24,24]
	   }
	 ]
	}`)
	if code != 1 {
		panic("set AutoMoveCoor http request err" + strconv.Itoa(code))
	}

	// 创建联盟成员
	//names := GenName(300)
	//wg.Add(300)
	//for _, v := range names {
	//	GenWonPlayerData(v, "", false, nil)
	//}
	//wg.Wait()

	// 3. 设置集群时间到开战期
	//code = WonPost("http://localhost:8080/request?account=sxf&rpc=20610", `{"sn":0,"guide":0,"trigger":"","str":"setClusterTime=2021-08-24 19:01:00"}`)
	//if code != 1 {
	//	panic("setClusterTime error" + strconv.Itoa(code))
	//}

	fmt.Println("success")
}

func StartMove(account string) {
	url := "http://localhost:8080/request?account=" + account + "&rpc="
	code := WonPost(url+"40006", `{"sn":0,"guide":0,"trigger":"","mapId":1,"coor":[24,24],"ft":42}`)
	if code != 1 {
		panic("enterWon http request err" + strconv.Itoa(code))
	}
	code = WonPost(url+"40006", `{"sn":0,"guide":0,"trigger":"","mapId":1,"coor":[24,24],"ft":43}`)
	if code != 1 {
		panic("enterWon http request err" + strconv.Itoa(code))
	}
	code = WonPost(url+"40006", `{"sn":0,"guide":0,"trigger":"","mapId":1,"coor":[24,24],"ft":44}`)
	if code != 1 {
		panic("enterWon http request err" + strconv.Itoa(code))
	}
}

func WonPost(url string, data string) int {
	jsonStr := []byte(data)
	resp, err := client.Post(url, ContentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	temp := respData{}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(result, &temp)
	return temp.Data.Code
}

// 生成问鼎中原玩家数据
func GenWonPlayerData(account string, guildName string, createGuild bool, houseCoor []int32) {
	//defer wg.Done()
	url := "http://localhost:8080/request?account=" + account + "&rpc="
	// 1.添加物品
	WonPost(url+"20610", `{"sn":0,"guide":0,"trigger":"","str":"addallitem=10"}`)

	// 2.godmode
	WonPost(url+"20610", `{"sn":0,"guide":0,"trigger":"","str":"godmode"}`)
	var code int
	// 3.创建或加入联盟
	if createGuild { // 3.创建联盟
		ok, err := CreateGuild(guildName, account)
		if err != nil || !ok {
			panic("create guild http request err")
		}
		// 升级联盟
		code = WonPost(url+"20610", `{"sn":0,"guide":0,"trigger":"","str":"guildLv=25"}`)
		if code != 1 {
			panic("guildLv error" + strconv.Itoa(code))
		}
	} else { //加入联盟
		// 随机加入联盟
		var guildId uint64

		if count < 65 {
			guildId = 777519105
		} else if count < 130 {
			guildId = 777519106
		} else if count < 195 {
			guildId = 777519107
		} else if count < 260 {
			guildId = 777519108
		} else if count < 320 {
			guildId = 777519109
		}
		count++
		ok, err := JoinGuild(guildId, account)
		if err != nil || !ok {
			fmt.Printf("join guild http request err [%v] ,ok [%v]", err, ok)
			fmt.Println()
		}
	}

	// 4.进入问鼎中原
	code = WonPost(url+"40000", `{"sn":0,"guide":0,"trigger":""}`)
	if code != 1 {
		panic("enterWon http request err" + strconv.Itoa(code))
	}

	// 5.进入地图
	code = WonPost(url+"40001", `{"sn":0,"guide":0,"trigger":"","mapId":1}`)
	if code != 1 {
		panic("enterMap http request err" + strconv.Itoa(code))
	}

	if createGuild {

		ok, err := WonAddHouse(account, 1, houseCoor)
		if err != nil || !ok {
			panic("addHouse http request err")
		}

		// 7.宣战城池
		code = WonPost(url+"40010", `{"sn":0,"guide":0,"trigger":"","mapId":1,"coor":[24,24]}`)
		if code != 1 {
			panic("declare grid http request err" + strconv.Itoa(code))
		}
	}

	// 8.保存布阵
	code = WonPost(url+"20500", `{"sn":0,"guide":0,"trigger":"","formation":[
   {
   "type":42,
   "CounsellorId":30008,
   "grid":[
       {
      "id":2,
      "commanderId":10066
      },
      {
      "id":5,
      "commanderId":10072
      },
     {
       "id": 6,
       "commanderId": 10069
     },
     {
       "id": 7,
       "commanderId": 10006
     },
     {
       "id":8,
       "commanderId":10016
     },
     {
       "id":10,
       "commanderId":10087
     },
     {
       "id": 11,
       "commanderId": 10007
     },
     {
       "id": 14,
       "commanderId": 10073
     }]
  },
   {
   "type":43,
   "CounsellorId":30011,
   "grid":[
     {
       "id":4,
       "commanderId":10501
     },
     {
       "id":6,
       "commanderId":10067
     },
     {
       "id": 7,
       "commanderId": 10054
     },
     {
       "id": 8,
       "commanderId": 10009
     },
     {
       "id":10,
       "commanderId":10070
     },
     {
       "id":11,
       "commanderId":10085
     },
     {
       "id": 12,
       "commanderId": 10090
     },
     {
       "id": 16,
       "commanderId": 10013
     }]
   },
   {
   "type":44,
   "CounsellorId":30013,
   "grid":[
     {
       "id":6,
       "commanderId":10004
     },
     {
       "id":7,
       "commanderId":10035
     },
     {
       "id": 8,
       "commanderId": 10061
     },
     {
       "id": 10,
       "commanderId": 10027
     },
     {
       "id":11,
       "commanderId":10003
     },
     {
       "id":12,
       "commanderId":10021
     },
     {
       "id": 14,
       "commanderId": 10050
     },
     {
       "id": 16,
       "commanderId": 10040
     }]
   }
 ]
}
`)
	if code != 1 {
		panic("save formation http request err" + strconv.Itoa(code))
	}

	code = WonPost(url+"40035", `{"sn":0,"guide":0,"trigger":"","mapId":1,
    "wonFtCoor":[
    {
      "ft": 42,
      "coor":[24,24]
    }
  ]
}`)
	if code != 1 {
		panic("set AutoMoveCoor http request err" + strconv.Itoa(code))
	}
	code = WonPost(url+"40035", `{"sn":0,"guide":0,"trigger":"","mapId":1,
    "wonFtCoor":[
    {
      "ft": 43,
      "coor":[24,24]
    }
  ]
}`)
	if code != 1 {
		panic("set AutoMoveCoor http request err" + strconv.Itoa(code))
	}

	code = WonPost(url+"40035", `{"sn":0,"guide":0,"trigger":"","mapId":1,
    "wonFtCoor":[
    {
      "ft": 44,
      "coor":[24,24]
    }
  ]
}`)
	if code != 1 {
		panic("set AutoMoveCoor http request err" + strconv.Itoa(code))
	}

}

func CreateGuild(guildName string, account string) (bool, error) {
	url := "http://localhost:8080/request?account=" + account + "&rpc="
	data := make(map[string]interface{})
	data["sn"] = 0
	data["guide"] = 0
	data["trigger"] = ""
	data["guildName"] = guildName
	data["declaration"] = "欢迎活跃人士加入"
	data["word"] = "吴"
	data["badge"] = 4
	data["flagSide"] = 1
	data["color"] = 3
	ok, err := Post(url+"21402", data)
	return ok, err
}

func JoinGuild(guildId uint64, account string) (bool, error) {
	url := "http://localhost:8080/request?account=" + account + "&rpc="
	data := make(map[string]interface{})
	data["sn"] = 0
	data["guide"] = 0
	data["trigger"] = ""
	data["guildId"] = guildId
	ok, err := Post(url+"21403", data)
	return ok, err
}

func WonAddHouse(account string, mapId uint32, coor []int32) (bool, error) {
	url := "http://localhost:8080/request?account=" + account + "&rpc="
	data := make(map[string]interface{})
	data["sn"] = 0
	data["guide"] = 0
	data["trigger"] = ""
	data["mapId"] = mapId
	data["coor"] = coor
	ok, err := Post(url+"40003", data)
	return ok, err
}

func Post(url string, data interface{}) (bool, error) {
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, ContentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	temp := respData{}
	result, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(result, &temp)
	if temp.Data.Code != 1 {
		return false, nil
	}

	return true, err
}

type respData struct {
	Message string `json:"message"`
	Data    rpData `json:"data"`
}

type rpData struct {
	Code int `json:"code"`
}

//随机生成玩家账号
func GetRandomString(l int) string {
	str := "0123456789acdehijklnopqrstuvwxyzACDEHIJKLNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	time.Sleep(1 * time.Microsecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenName(len int) []string {
	names := make([]string, 0, len)
	for i := 1; i <= len; i++ {
		s := "suxf" + strconv.Itoa(i)
		names = append(names, s)
	}
	return names
}

//func reqAccount(url string, count int) ([]string, error) {
//
//	resp, err := client.Get(url + "?count=" + strconv.Itoa(count))
//	if nil != err {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	b1, err := ioutil.ReadAll(resp.Body)
//	if nil != err {
//		return nil, err
//	}
//
//	var ret []string
//	if err := json.Unmarshal(b1, &ret); nil != err {
//		return nil, err
//	}
//	return ret, nil
//}
//
//// 获取线上账号
//func GetOnlineAccount() ([]string, error) {
//	url := "http://127.0.0.1:10000/account"
//	return reqAccount(url, 500)
//}
