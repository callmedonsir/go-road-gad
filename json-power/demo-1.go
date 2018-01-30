package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type rpcNode struct {
	IP net.IP
}

type neighbors struct {
	Nodes []rpcNode
}

func main() {
	nodes := neighbors{Nodes: []rpcNode{rpcNode{IP: []byte{127, 0, 0, 1}}, {IP: []byte{127, 0, 0, 2}}}}
	node := make([][]rpcNode, 0)
	node = append(node, nodes.Nodes)
	a, _ := json.Marshal(node)
	str := strings.Split(strings.Split(strings.Split(string(a), ",")[1], ":")[1], "\"")
	fmt.Println(str[1])
}
