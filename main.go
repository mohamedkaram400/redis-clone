package main

import (
	"fmt"
	// "os"
	// "bufio"
	// "strconv"
	// "io"
	// "strings"

	"net"
)


func main() {
	fmt.Println("Listening on port :6379")

	// Create new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// input := "$5\r\nAhmed\r\n"
	// reader := bufio.NewReader(strings.NewReader(input))

	// b, _ := reader.ReadByte()

	// if b != '$' {
	// 	fmt.Println("Invaild type, expecting bulk strings only")
	// 		os.Exit(1)
	// }

	// size, _ := reader.ReadByte()

	// strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// // consume /r/n
	// reader.ReadByte()
	// reader.ReadByte()

	// name := make([]byte, strSize)
	// reader.Read(name)

	// fmt.Println(string(name))

	for {
		reps := NewResp(conn)
		value, err := reps.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))

		// buf := make([]byte, 1024)

		// // Read message from client
		// _, err := conn.Read(buf)
		// if err != nil {
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	fmt.Println("error reading from cleint: ", err.Error())
		// 	os.Exit(1)
		// }

		// // Ignore request and send back a PONG
		// conn.Write([]byte("+OK\r\n"))
	}
}
