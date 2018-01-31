package main

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"sync"
	"time"
	"os/exec"
)

func main(){
	sigRecv1 := make(chan os.Signal,1)
	sigs1:=[]os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRevc1]\n",sigs1)
	signal.Notify(sigRecv1,sigs1...)

	sigRecv2 :=make(chan os.Signal,1)
	sigs2 :=[]os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRevc2]\n",sigs2)
	signal.Notify(sigRecv2,sigs2...)
	wg:=sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for sig:=range sigRecv1{
			fmt.Printf("received a singal from sigRecv1 :%s\n",sig)
		}
		fmt.Println("end [sigRecv]")
		wg.Done()
	}()

	go func() {
		for sig:=range sigRecv2{
			fmt.Println("reveived a singel from sigRecv2",sig)

		}
		fmt.Println("end [sigRecv2]")
		wg.Done()
	}()
	fmt.Println("wait for 2 seconds")
	time.Sleep(time.Second*2)

	fmt.Println("stop notification....")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Println("done sigRecv1")
	wg.Wait()


	cmds:=[]*exec.Cmd{
		exec.Command("ps","aus"),
		//ps aus
		exec.Command("grep","signal"),
		exec.Command("grep","-v","grep"),
		exec.Command("grep","-v","go run"),
		exec.Command("awk","{print $2}"),
	}

	output,err:=runCmds(cmds)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%#v",output)
}


func runCmds(cmds []*exec.Cmd)([]string,error){
	str := make([]string,0)
	for i:=0;i<len(cmds);i++{
		output,err:=cmds[i].StdoutPipe()
		if err!=nil{
			return nil,err
		}
		cmds[i].Start()
		outPut0 :=make([]byte,100)
		n,err:=output.Read(outPut0)
		if err!=nil{
			return nil,err
		}
		str[i]=string(outPut0[:n])
	}
	return str,nil
}