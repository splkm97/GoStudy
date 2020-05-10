package main

import "fmt"

const (
    Big   = 1 << 100	// 2^100
    Small = Big >> 99	// (2^100) / (2^99) = 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small)) // 21
    // 아래 두줄 실행 안됨.. -> int범위 벗어난 상수를 fmt.Println()으로 출력할 수 없다.
    // fmt.Println(Big)
    // fmt.Println(needInt(Big))
    fmt.Println(needFloat(Small)) // 0.2
    fmt.Println(needFloat(Big)) // 1.2676506002282295e+29 -> 100제곱이어야 하지만 너무 커서 잘린 듯..
}
