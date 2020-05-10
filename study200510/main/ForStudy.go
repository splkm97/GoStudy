package main

import "fmt"

func main() {
    // 일반적인 for 사용법
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
    
    // while 을 대체하는 for
    for = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
    
    
    /*
    -- 무한루프: 조건문 없이 for 반복 --
    for {
    }
    */
}
