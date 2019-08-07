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
	// file.Read()可以读取一个小文件到大的bytgo e slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误 例如：EOF
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
