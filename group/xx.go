package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
	"runtime"
)

type down struct {
	Lock sync.Mutex
	ch   chan string
	wg   sync.WaitGroup
}

func main() {
	done := down{ch: make(chan string, 100000)}
	runtime.GOMAXPROCS(runtime.NumCPU())
	for j := 0; j < 100000; j++ {
		done.wg.Add(1)
		done.ch <- strconv.Itoa(j)
	}
	defer close(done.ch)
	var fileName = "./b.txt"
	var file *os.File
	var err1 error
	if check(fileName) {
		file, err1 = os.OpenFile(fileName, os.O_APPEND, 0777)
	} else {
		file, err1 = os.Create(fileName)
	}
	if err1 != nil {
		fmt.Println("some thing wrong ", err1)
	}
	defer file.Close()
	time1:=time.Now()
	go func() {
		for ch:=range done.ch{
			done.Lock.Lock()
			done.wg.Done()
			_, err := io.WriteString(file, ch)
			if err != nil {
				fmt.Println(err)
			}
			done.Lock.Unlock()
		}

	}()
	fmt.Println(time.Since(time1))
	done.wg.Wait()


}

func check(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
