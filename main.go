package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// big endian data (1000) of 32 bit
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// endian tranform
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
