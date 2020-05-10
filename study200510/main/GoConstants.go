package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "안녕"
    // World = "세상" // 상수 변경시 
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
    
    const char1 = 'A'
    fmt.Println(char1)	// 'A'의 아스키코드값 65 출력 (???)
    fmt.Printf("%c", char1)	// A 출력
}
