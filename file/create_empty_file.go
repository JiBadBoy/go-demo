package main

import (
	"log"
	"os"
)



func main()  {
	var (
		newFile *os.File
		err error
	)
	//	创建空文件，如果存在将会覆盖清空
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

