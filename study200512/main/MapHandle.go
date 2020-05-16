package main

import "fmt"

func main() {
    // 맵을 만드는 make
    m := make(map[string]int)

    // 맵에 key-value값 지정
    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])
	
    // 맵의 원소 삭제
    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    // 맵에 키("Answer")가 있는지 확인
    // 있다면 ok -> true, v -> key 값에 대응되는 value
    // 없다면 ok -> false, v -> 0 (zero value)
    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
