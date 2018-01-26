package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Println("you may have to import like go run xxx.go YOUR_IMPORT")
		return
	}
	log.Println("start get information from tcp by port 8888")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("some error in Dial ", err)
		return
	}
	defer conn.Close()
	data:=os.Args[1]

	n, err := conn.Write([]byte(data))
	if err != nil {
		log.Println("some error happend in write ", err)
		return
	}
	log.Printf("send complement send %d words\n",n)
}
