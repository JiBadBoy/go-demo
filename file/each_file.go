package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	currentPath, _ := os.Getwd()
	// 获取上级目录
	parentDir := getParentDirectory(currentPath)

	//walkByFilePath(getParentDirectory(currentPath))
	walkByIoutil(parentDir, ".go")
}

// filepath包遍历
func walkByFilePath(path string)  {
	filepath.Walk(path, func(pathSub string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(pathSub)
		return nil
	})
}


// ioutil包遍历
func walkByIoutil(path string, suffix string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	// 系统分隔符
	pathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			walkByIoutil(path + pathSep + fi.Name(), suffix)
			//walkByIoutil(filepath.Join(path, fi.Name()), suffix) // 两种方式拼接
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			fmt.Println(path + pathSep + fi.Name())
		}
	}
	return nil
}


func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, string(os.PathSeparator)))
}