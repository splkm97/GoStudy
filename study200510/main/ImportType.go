package main

// 여러 개의 패키지 import 방식 1
import (
  "fmt"
  "math"
)

/*
// 여러 개의 패키지 import 방식 2
import "fmt"
import "math"
*/

func main() {
  fmt.Println("Happy", math.Pi, "Day")
}
