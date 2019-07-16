package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// 方法接收者 出现在 func 关键字和方法名之间
func (v *Vertex) Abs() float64 { // 注意的方法接受者是指针类型
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
