package main

import (
	"fmt"
	"reflect"
)

type order2 struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery2(q interface{}) {
	val := reflect.ValueOf(q)
	if val.Kind() != reflect.Ptr {
		fmt.Println("unsupported type")
		return
	}
	t := reflect.TypeOf(q).Name()
	typ := reflect.Indirect(val).Type()
	a := typ.Name()
	fmt.Println(a)
	query := fmt.Sprintf("insert into %s values(", t)
	v := reflect.ValueOf(q)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
			} else {
				query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
			} else {
				query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
			}
		default:
			fmt.Println("Unsupported type")
			return
		}
	}
	query = fmt.Sprintf("%s)", query)
	fmt.Println(query)
}

func main() {
	o := &order2{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery2(e)
	i := 90
	createQuery2(i)
}
