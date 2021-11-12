package main

import "fmt"

type ISuperDraw interface {
	Draw(count int32)
	TakeBoxReward(boxId int32)
	CreateClientData()
	OnCrossDay()
	OnClose(bool)
}

var _ ISuperDraw = (*UrDraw)(nil)

type UrDraw struct {
	He int
	Me string
}

func NewUrDraw(a int, b string) ISuperDraw {
	return &UrDraw{
		He: a,
		Me: b,
	}
}

func (u UrDraw) Draw(count int32) {
	panic("implement me")
}

func (u UrDraw) TakeBoxReward(boxId int32) {
	panic("implement me")
}

func (u UrDraw) CreateClientData() {
	panic("implement me")
}

func (u UrDraw) OnCrossDay() {
	panic("implement me")
}

func (u UrDraw) OnClose(b bool) {
	panic("implement me")
}

func main() {
	temp := make(map[int]ISuperDraw)
	temp[1] = NewUrDraw(98, "sxf")
	fmt.Println(temp)
}
