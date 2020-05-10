package main

// 타입 확인 위한 reflect 패키지 import
import "fmt"
import "reflect"

var x, y, z int = 1, 2, 3
var c, python, java = true, 'F', "no!"

func main() {
    fmt.Println(x, y, z, c, python, java)
    fmt.Println("Type of x: ", reflect.TypeOf(x))
    fmt.Println("Type of y: ", reflect.TypeOf(y))
    fmt.Println("Type of z: ", reflect.TypeOf(z))
    fmt.Println("Type of c: ", reflect.TypeOf(c))
    fmt.Println("Type of python: ", reflect.TypeOf(python))
    fmt.Println("Type of java: ", reflect.TypeOf(java))
    // 실행 결과 x, y, z, c, python, java 는 각각
    // int, int, int, bool, int32, string이었다.
    // char -> int32인건가?
}
