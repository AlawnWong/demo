package main

import "fmt"

func worker(c chan int) {
	//从channel中去读数据
	num := <-c
	fmt.Println("foo recv channel ", num)
}

func main() {
	//创建一个channel
	c := make(chan int)

	go worker(c)

	//main协程 向一个channel中写数据
	c <- 1

	fmt.Println("send 1 -> channel over")
}
