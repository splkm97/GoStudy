package main

import "fmt"

// 세제곱근을 구하는 함수
func Cbrt(x complex128) complex128 {
    var z complex128 = 1.
    for i := 0; i < 10; i ++ {
        z = z - (z * z * z - x) / (3 * z * z)
    }
    
    return z
}

func main() {
    fmt.Println(Cbrt(2))
}
