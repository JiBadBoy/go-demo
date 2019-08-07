package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// 读取文件到byte slice中   和 ReadAll一样 可以看源码
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}
