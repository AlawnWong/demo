package main

import (
	"math"
	"fmt"
)

type VertexP struct {
	X, Y float64
}

func (v *VertexP) ScaleP(f float64) {
	// 如果这里不是指针类型而是值类型，而传入的是原始结构体的一个复制后的副本
	// 修改副本是不会影响原始结构体的数据的
	// 而且会有复制结构体的开销
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *VertexP) AbsP() float64 {
	// 如果这里不是指针类型而是值类型，而传入的是原始结构体的一个复制后的副本
	// 因为是只读，读取原始结构体和副本，数值都相同，不影响计算结果
	// 但是会有复制结构体的开销
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &VertexP{3, 4}
	v.ScaleP(5)
	fmt.Println(v, v.AbsP())
}
