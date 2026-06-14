package main

import (
	"bufio"
	"os"
	"sync"
	"time"

	"github.com/mohamedkaram400/redis-clone/resp"
)

type AOF struct {
	file *os.File
	rd bufio.Reader
	mu sync.Mutex
}

func NewAOF(path string) (*AOF, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	aof := &AOF{
		file: f,
		rd: *bufio.NewReader(f),
	}

	// Start a goroutine to sync AOF to disk every 1 second
	go func ()  {
		for {
			aof.mu.Lock()
			aof.file.Sync()
			aof.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return aof, nil
}

func (aof *AOF) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	return aof.Close()
}

func (aof *AOF) Write(value resp.Value) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return err
	}

	return nil
}

// func (*aof *AOF) Read() {

// }