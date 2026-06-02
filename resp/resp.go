package resp

import (
	"bufio"
	"io"
)


const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)


type Value struct {
	Typ string
	Str string
	Num int
	Bulk string
	Array []Value
}


type Resp struct {
	reader *bufio.Reader
}

type Writer struct {
	writer io.Writer
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}


// 🔥 The full flow (very important)
// 1. Client (terminal)
// You type:
// PING key

// 2. What is ACTUALLY sent (hidden from you)
// *2\r\n
// $4\r\nPING\r\n
// $3\r\nkey\r\n
// 👉 This is RESP protocol