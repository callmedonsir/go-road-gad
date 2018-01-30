package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printClientLog(id, "Dial Error %s", err)
		return
	}
	defer conn.Close()

	printClientLog(id, "connect to server .(remote address :%s , local address %s)",
		conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(time.Millisecond * 5))
	for i := 0; i < requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "write Error %s", err)
			continue
		}
		printClientLog(id, "send request (written %d bytes) : %d.", n, req)
	}
}
