package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    // 몇 번째 수열인지 나타내는 cnt
    // 현재 수열의 값 cur
    // 이전 수열의 값 prev
    // swap을 위한 temp
    cnt := 0
    cur := 1
    prev := 1
    temp := 1
    
    return func() int {
        // 매 호출마다 cnt 증가
        cnt ++
        // 첫 번째, 두 번째는 1을 리턴
        if cnt == 1 || cnt == 2 {
            return 1
        } else {
            // 그 이외에는 cur을 prev로 바꾸고
            // cur에 prev + cur 를 할당
            temp = prev
            prev = cur
            cur += temp
            // 현재 피보나치 수를 리턴
            return cur
        }
	}
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
