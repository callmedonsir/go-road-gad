package main

import (
	"os"
	"fmt"
	"io"
	"strconv"
	"time"
)

func main(){

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
	for i:=0;i<100000;i++{
		io.WriteString(file,strconv.Itoa(i))
	}
	fmt.Println(time.Since(time1))
}
func check(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}