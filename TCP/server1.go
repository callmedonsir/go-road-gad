package main

import (
	"net"
	"log"
)

func main(){
	c,err:=net.Listen("tcp",":8888")
	if err!=nil{
		log.Println("listen err",err)
		return



	}


	for{
		get,err:=c.Accept()
		if err!=nil{
			log.Println("accept err",err)
			break
		}
		log.Println("start a new connection")
		go handleConn(get)
	}

}
func handleConn(get net.Conn){
	defer get.Close()
	for{
		var str = make([]byte,10)
		log.Println("start read from conn")
		n,err:=get.Read(str)
		if err!=nil{

			log.Println("read err",err)
			return
		}
		log.Printf("read %d words , these is %s",n,string(str[:n]))

	}
}

