package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outPut1 bytes.Buffer
	cmd1.Stdout = &outPut1
	if err := cmd1.Start(); err != nil {
		fmt.Println("cmd1 can not be startup", err)
		return
	}
	//cmd1的wait方法会让cmd1阻塞 直到cmd1完全运行结束
	if err := cmd1.Wait(); err != nil {
		fmt.Println("wait err:", err)
		return
	}
	cmd2.Stdin = &outPut1
	var outPut2 bytes.Buffer
	cmd2.Stdout = &outPut2

	if err := cmd2.Stdout; err != nil {
		fmt.Println("cmd2 start err", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Println("cmd2 wait err", err)
		return
	}

	fmt.Printf("%s\n", outPut2)
}
