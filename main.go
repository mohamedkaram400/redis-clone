package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/mohamedkaram400/redis-clone/resp"
)


func main() {
	fmt.Println("Listening on port :6379")

	// Create new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	aof, err := NewAOF("database.aof")
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer aof.Close()

	aof.Read(func(value resp.Value) {
		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return 
		}
		handler(args)
	})

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection after method executed
	defer conn.Close()


	for {
		reps := resp.NewResp(conn)
		value, err := reps.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]
		
		writer := resp.NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: "ERR unknown command"})
			continue
		}

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}
		result := handler(args)
		writer.Write(result)
	}
}
