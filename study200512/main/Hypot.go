package main

import (
    "fmt"
    "math"
)

func main() {
    // 함수를 hypot이라는 변수에 할당
    // 피타고라스 정리
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(3, 4))
}
