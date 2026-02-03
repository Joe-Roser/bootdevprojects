package main

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	if err != nil {
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		utils.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		in, err := reader.ReadBytes('\n')
		if err != nil {
			utils.Fatal(err)
		}

		conn.Write(in)
	}
}
