package main

import (
    "math"
    "fmt"
)

// 에러 타입 설정
type ErrNegativeSqrt float64

// 에러 메소드 구현
func (e ErrNegativeSqrt) Error() string {
    ans := fmt.Sprintf("cannot Sqrt negative number: %f", e)
    return ans
}

// 에러 메소드 처리
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    ans := math.Sqrt(f)
    return ans, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
