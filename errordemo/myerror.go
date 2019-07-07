package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f:%s", e.radius, e.err)
}

func (e *areaError) IsRadiusNagative() bool {
	return e.radius < 0

}
func circleAreaMy(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"Radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	s, err := circleAreaMy(-20)
	if err != nil {
		//将错误转换为具体的类型
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
