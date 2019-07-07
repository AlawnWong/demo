package main

import (
	"fmt"
	"os"
)

func main() {
	//试图打开一个并不存在的文件，这将会返回一个error
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err) //no such file or directory
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
