package main

import (
    "fmt"
    "math"
)

// 인터페이스 정의: Abs() 라는 메소드를 정의해야함
type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser
    // a = v  // a Vertex, does NOT
    // implement Abser

    fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloat 의 Abs() 메소드 정의
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

// Vertex 의 Abs() 메소드 정의
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
