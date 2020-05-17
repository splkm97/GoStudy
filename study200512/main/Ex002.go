package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // 이중 슬라이스 생성
    s := make([][]uint8, dy)
    for i := 0; i < dy; i ++ {
        // 슬라이스 각 요소마다 슬라이스 생성
        s[i] = make([]uint8, dx)
        for j := 0; j < dx; j ++ {
            // 각 슬라이스 자리에 숫자 표현
            s[i][j]  = uint8((i+j)/2)
        }
    }
    return s
}

func main() {
   /*
   pic 패키지 -> import "golang.org/x/tour/pic"
   pic.Show() -> [][]uint8 슬라이스를 리턴하는 함수를 인자로 전달하면
               그 함수에 맞는 패턴을 흑백으로 콘솔에 띄움
   pic.ShowImage() -> 이미지를 콘솔에 띄움
   */
    pic.Show(Pic)
}
