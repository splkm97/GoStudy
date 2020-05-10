package main

import "fmt"

// 다음과 같이 2개 이상의 매개변수가 같은 타입일때
// 한번에 타입 지정이 가능하다.
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
