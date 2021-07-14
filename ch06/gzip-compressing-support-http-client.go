package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// retry loop
	for {
		var err error
		if conn == nil {
			// Dial
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access %d\n", current)
		}

		// make POST request
		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Endcoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		// read from server. retry when timeout
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// display result
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		// end if all sent
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
