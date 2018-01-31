package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"bufio"
)

func main(){
	cmd0:=exec.Command("echo","-n","my first comes from golang")
	//对应shell 命令为 echo -n "my first comes from golang"
	stdout0,err:=cmd0.StdoutPipe()
	if err!=nil{
		fmt.Printf("ERROR: can not botain the stdout pipe for command no.0;%s\n",err)
		return
	}
	if err:=cmd0.Start();err!=nil{
		fmt.Printf("ERROR: the command no.0 can not be start up :%s \n",err)
		return
	}

//第一种获取方式
	outPut0 :=make([]byte,100)
	n,err:=stdout0.Read(outPut0)
	if err!=nil{
		fmt.Printf("can not read date from pipe :%s\n",err)
		return
	}
	fmt.Printf("data is %s\n",outPut0[:n])

//第二种获取方式
	var outPutBuffer0 bytes.Buffer
	for{
		tempOutPut:=make([]byte,5)
		n,err:=stdout0.Read(tempOutPut)
		if err!=nil{
			if err==io.EOF {
				break
			}else{
				panic(err)
				return
			}
		}
		if n>0 {
			outPutBuffer0.Write(tempOutPut[:n])
		}
	}
	fmt.Println(outPutBuffer0.String())

	//第三中获取方式
	outPutBuffer1:=bufio.NewReader(stdout0)
	outPut1,_,err:=outPutBuffer1.ReadLine()
	if err!=nil{
		panic(err)
		return
	}
	fmt.Printf("%s\n",string(outPut1))
}
