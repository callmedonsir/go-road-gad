package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	//作为一个表示数据边界的单字节字符
	DELIMITER = '\t'
)

func serverGo() {
	defer wg.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("listen Error: %s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listen for the server.(local address:%s)", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept err", err)
		}
		printServerLog("Established a connection with a client application,(remote address %s)",
			conn.RemoteAddr())
		go handleConn(conn)
	}
}

func printServerLog(format string, args ...interface{}) {
	printLog("server", 0, format, args...)
}

func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]:%s", role, sn, fmt.Sprintf(format, args))
}

func printClientLog(sn int, format string, args ...interface{}) {
	printLog("client", sn, format, args)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		strReq, err := read(conn)
		if err != nil {
			if err != io.EOF {
				printServerLog("the connection is closed by another side.")
			} else {
				printServerLog("read error %s", err)
			}
			break
			printServerLog("received request %s", strReq)
		}

		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("send error message(written %d bytes): %s.", n, err)
			continue
		}
		floatResp := int(math.Cbrt(float64(intReq)))
		respMsg := fmt.Sprintf("the cube root of %d is %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("write error : %s.", err)
		}
		printServerLog("Sent response (written %d bytes):%s", n, respMsg)
	}

}

func strToInt32(strReq string) (int32, error) {
	intStrReq, err := strconv.Atoi(strReq)
	return int32(intStrReq), err
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}
