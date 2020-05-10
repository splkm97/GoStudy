package main

import (
  "fmt"
  "math"
)

func main() {
  // fmt.Println(math.pi) 
  // 소문자로 시작된 pi: export 되어있지 않다-> error 발생
  fmt.Println(math.Pi)
  // 대문자로 시작된 Pi: export 되어있다
}
