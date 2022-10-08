package main

import (
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	engine, err := xorm.NewEngine("mysql", "root:ld&FsNtKjsZ$1Io^@tcp(47.98.211.160:3306)/hualing?charset=utf8")
	// engine2, err := xorm.NewEngine("mysql", "root:123456@tcp(119.45.254.67:3306)/chat?charset=utf8")
	if err != nil {
		fmt.Println("connect mysql is failed, err:", err)
	}
	// dd := make([]interface{}, 0)
	total, err := engine.Where("suid = ?", 3765549).Count(&NftOrder{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
}

const hex = 5 // 5进制

// AddFiveHex 五进制，加一算法
func AddFiveHex(initVal int64) (string, error) {
	initStr := strconv.Itoa(int(initVal))
	lastValueStr := string(initStr[len(initStr)-1])
	lastValueInt, err := strconv.Atoi(lastValueStr)
	if err != nil {
		return "", err
	}
	endInt := lastValueInt + 1
	ts := modifyStrValue(initStr, len(initStr)-1, endInt)
	if endInt <= hex {
		return ts, nil
	}

	// 发生了进位
	for i := len(ts) - 1; i >= 0; i-- {
		curStr := string(ts[i])
		curInt, err := strconv.Atoi(curStr)
		if err != nil {
			panic(err)
		}
		if curInt <= hex {
			return ts, nil
		} else {
			// 防止下面的nextStr数组越界
			if i == 0 {
				return "", errors.New("已经满值了")
			}
		}
		nextStr := string(ts[i-1])
		nextInt, err := strconv.Atoi(nextStr)
		if err != nil {
			panic(err)
		}
		m1 := modifyStrValue(ts, i, 0)
		ts = modifyStrValue(m1, i-1, nextInt+1)
	}
	return ts, nil
}

// 将字符串str的index位修改成finVal值
func modifyStrValue(str string, index int, finVal int) string {
	c := []byte(str)
	temp := strconv.Itoa(finVal)
	c[index] = []byte(temp)[0]
	return string(c)
}

func CalSerialNumber(version int64, number int64) string {
	// 初始长度5位
	initLen := 5
	verStr := strconv.Itoa(int(version))
	numStr := strconv.Itoa(int(number))
	if len(numStr) < initLen {
		var temp string
		for i := 0; i < initLen-len(numStr); i++ {
			temp += "0"
		}
		return verStr + temp + numStr
	} else {
		return verStr + numStr
	}
}

func GenerateToken(quality int64, tokenNumber int64) string {
	qstr := strconv.Itoa(int(quality))
	str := strconv.Itoa(int(tokenNumber))
	var res string
	for _, v := range str {
		temp := "0" + string(v)
		res += temp
	}
	return qstr + res
}
