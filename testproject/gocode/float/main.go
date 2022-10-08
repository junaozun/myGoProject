package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// var number_one = decimal.NewFromFloat(0.59)
	// var number_two = decimal.NewFromFloat(0.12587)
	//
	// var number_three = number_one.Add(number_two) // 计算两个浮点数相加
	// var number_four = number_one.Sub(number_two)  // 计算两个浮点数相减
	// var number_five = number_one.Mul(number_two)  // 计算两个浮点数相乘
	// var number_six = number_five.Div(number_two)  // 计算两个浮点数相除
	//
	// fmt.Println(number_one, number_two, number_three, number_four, number_five, number_six)
	//
	// var number_seven = decimal.NewFromFloat(3.6987)
	// var number_eight = number_seven.Round(2)   // 保留小数点后两位（四舍五入）
	// var number_nine = number_seven.Truncate(2) // 保留小数点后两位
	// v, ok := number_nine.Float64()
	// fmt.Println(v, ok)
	// fmt.Println(number_seven, number_eight, number_nine)
	var a float64 = 34.986432
	var b float64 = 89.7534523256
	c := Float64Mul(a, b, false, 0)
	fmt.Println(c)
}

// Float64Add num1+num2
func Float64Add(num1 float64, num2 float64, isRound bool, floatCount int32) float64 {
	a := decimal.NewFromFloat(num1)
	b := decimal.NewFromFloat(num2)
	c := a.Add(b)
	var d decimal.Decimal
	if isRound {
		d = c.Round(floatCount)
	} else {
		d = c.Truncate(floatCount)
	}
	return d.InexactFloat64()
}

// Float64Sub num1-num2
func Float64Sub(num1 float64, num2 float64, isRound bool, floatCount int32) float64 {
	a := decimal.NewFromFloat(num1)
	b := decimal.NewFromFloat(num2)
	c := a.Sub(b)
	var d decimal.Decimal
	if isRound {
		d = c.Round(floatCount)
	} else {
		d = c.Truncate(floatCount)
	}
	return d.InexactFloat64()
}

// Float64Mul num1*num2
func Float64Mul(num1 float64, num2 float64, isRound bool, floatCount int32) float64 {
	a := decimal.NewFromFloat(num1)
	b := decimal.NewFromFloat(num2)
	c := a.Mul(b)
	var d decimal.Decimal
	if isRound {
		d = c.Round(floatCount)
	} else {
		d = c.Truncate(floatCount)
	}
	return d.InexactFloat64()
}

// Float64Div num1/num2
func Float64Div(num1 float64, num2 float64, isRound bool, floatCount int32) float64 {
	a := decimal.NewFromFloat(num1)
	b := decimal.NewFromFloat(num2)
	c := a.Div(b)
	var d decimal.Decimal
	if isRound {
		d = c.Round(floatCount)
	} else {
		d = c.Truncate(floatCount)
	}
	return d.InexactFloat64()
}
