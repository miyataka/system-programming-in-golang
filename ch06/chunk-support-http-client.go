package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	var conn net.Conn = nil
	var err error
	// Dial
	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	// make POST request
	request, err := http.NewRequest("POST", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	// read from server. retry when timeout
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	// display result
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 ||
		response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		line := make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf(" %d bytes: %s\n", size, string(line))
	}

	conn.Close()
}
