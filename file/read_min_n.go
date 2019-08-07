package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	byteSlice := make([]byte, 512)
	minBytes := 20
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
