package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("  file name: %v\n", info.Name())
	fmt.Printf("  size: %v\n", info.Size())
	fmt.Printf("  modified time: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  directory?: %v\n", info.Mode().IsDir())
	fmt.Printf("  regular file?: %v\n", info.Mode().IsRegular())
	fmt.Printf("  Unix file permission bit:%v\n", info.Mode().Perm())
	fmt.Printf("  mode text: %v\n", info.Mode().String())
}
