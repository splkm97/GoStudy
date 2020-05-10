package main

import "fmt"

          // 반환형에 x, y라는 이름 부여
func split(sum int) (x, y int) {
    // x, y라는 이름으로 반환값 변수처럼 사용
    x = sum * 4 / 9
    y = sum - x
    // 반환
    return
}

func main() {
    fmt.Println(split(17))
}
