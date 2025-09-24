package main

import (
	"bytes"
	"fmt"
	"io"
	"main/cmd/utils"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		utils.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			println(err)
			break
		}
		fmt.Println("Connection accepted")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\n", line)
		}
		fmt.Println()
	}

	listener.Close()
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		buf := make([]byte, 8)

		var current_line string
		for {
			n, err := f.Read(buf)

			if err != nil {
				break
			}

			read_in := buf[:n]
			if i := bytes.IndexByte(buf, '\n'); i != -1 {
				current_line += string(read_in[:i])
				out <- current_line
				current_line = string(read_in[i+1:])
			} else {
				current_line += string(read_in)
			}
		}

		if len(current_line) > 0 {
			out <- current_line
		}
	}()

	return out
}
