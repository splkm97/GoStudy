package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    const small = 0.000000000000001;
    z := 1.0
    for i:= 0; i < 10; i ++ {
        z = z - (z*z - x) / (2*z)
    }
    fmt.Println("10번 반복한 값: ", z)
    prevZ := z
    for {
        prevZ = z
        z = z - (z*z - x) / (2*z)
        if prevZ > z-small && prevZ < z+small {
            break
        }
    }
    
    return z
}

func main() {
    const small = 0.000000000000001;
    num := 10.0
    result := Sqrt(num)
    fmt.Printf("%g", small)
    fmt.Println("보다 작은 오차 발생시까지 반복한 값: ",result)
    fmt.Println("math.Sqrt(", num, ") 결과값: ", math.Sqrt(num))
}
