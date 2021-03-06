# 07 함수 (1)

C/C++/Java 와 다르게 함수 매개변수 타입을 뒤에 둔다

-> 왼쪽에서 오른쪽으로 읽을 때 편하다

-> 함수 add는 x정수 y정수를 매개변수로 하고 int를 반환한다.

-> 잘 이해가... 질문이 필요할 것 같다.

```go
int add(int x, int y) {
  return x+y;
}

Go:
func add(x int, y int) int {
  return x+y
}
```

# 08 함수 (2)

둘 이상의 매개변수가 같은 타입일때

동시에 타입 지정이 가능하다.
```go
func add(x int, y int) int {
    return x + y
}
```
or
```go
func add(x, y int) int {
    return x + y
}
```

# 09 여러 개의 결과 (Multiple results)

여러 개의 리턴값을 반환 가능하다.

```go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

# 10 이름이 정해진 결과 (Named results)

반환값에 이름을 붙여 변수처럼 사용 가능
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

# 11 변수 (Variables)

var 을 이용해 변수 선언

선언문 마지막에 타입 지정
```go
var x, y, z int
var a, b, c bool
```

# 12 변수 초기화

변수를 선언과 동시에 초기화 가능하다.

이 때 따로 타입 지정이 필요 없기도 하다.
```go
var x, y, z int = 1, 2, 3
var c, python, java = true, 'F', "no!"
Type of x:  int
Type of y:  int
Type of z:  int
Type of c:  bool
Type of python:  int32
Type of java:  string
```
