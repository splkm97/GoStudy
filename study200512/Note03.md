# 53, 54 인터페이스 (Interface)

인터페이스는 메소드들의 집합이다.

             < InterfaceExam.go >
```go
package main

import (
    "fmt"
    "math"
)

// 인터페이스 정의: Abs() 라는 메소드를 정의해야함
type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser
    // a = v  // a Vertex, does NOT
    // implement Abser

    fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloat 의 Abs() 메소드 정의
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

// Vertex 의 Abs() 메소드 정의
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

# 55 에러 (Error)

에러는 string 형을 반환하는 하나의 메소드 Error 로 구성된 내장 인터페이스 타입 error 를 구현한 것입니다.

fmt패키지의 출력 루틴들은 에러 출력을 요청받았을때 자동으로 Error() 메소드를 호출합니다.

```go
package main

import (
    "fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
```

# 56 연습: 에러

Ex006.go 로 풀었습니다.

