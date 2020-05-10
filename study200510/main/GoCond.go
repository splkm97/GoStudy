package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        // 허수 표현을 위해 i를 추가한 string
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
    // v 를 이용해 조건문 기능 추가
        return v
    } else {
    // else 블록 안에서도 v 를 이용할 수 잇다.
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // can't use v here, though
    return lim
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
    
    fmt.Println("-----------------")
    
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
