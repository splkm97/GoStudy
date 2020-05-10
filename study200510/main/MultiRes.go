package main

import "fmt"

                    // 두개 이상의 return 반환
func swap(x, y string) (string, string) {
    // 리턴함.
    return y, x
}

func main() {
    // 두개 이상의 return 받기
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
