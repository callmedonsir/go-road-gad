package main

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	go serverGo()
	time.Sleep(time.Millisecond * 500)
	go clientGo(1)
	wg.Wait()
}
