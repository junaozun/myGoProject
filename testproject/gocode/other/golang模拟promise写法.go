package main

import (
	"fmt"
	"sync"
	"time"
)

type Resolve func(rsp interface{})
type Reject func(err error)
type PromiseFunc func(resolve Resolve, reject Reject)

type Promise struct {
	f       PromiseFunc
	resolve Resolve
	reject  Reject
	wg      sync.WaitGroup
}

func NewPromise(f PromiseFunc) *Promise {
	return &Promise{f: f}
}

func (t *Promise) then(resolve Resolve) *Promise {
	t.resolve = resolve
	return t
}

func (t *Promise) catch(reject Reject) *Promise {
	t.reject = reject
	return t
}

func (t *Promise) done() {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.f(t.resolve, t.reject)
	}()
	t.wg.Wait()
}

func main() {
	NewPromise(func(resolve Resolve, reject Reject) {
		time.Sleep(time.Second)
		if time.Now().Unix()%2 == 0 {
			resolve("ok")
		} else {
			reject(fmt.Errorf("my error"))
		}
	}).then(func(rsp interface{}) {
		fmt.Println(rsp)
	}).catch(func(err error) {
		fmt.Println(err)
	}).done()
}
