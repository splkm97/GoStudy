# 41 연습: 맵

Ex003.go 로 풀었습니다.

# 42 함수 값 (Function values)

함수를 변수에 할당 가능하다.
```go
func main() {
    // 함수를 hypot이라는 변수에 할당
    // 피타고라스 정리
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(3, 4))
}
```

# 43 함수 클로져 (Function closures)

함수는 클로저(Full closures)입니다.
```go
/** 클로저 **/
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 아래 세 라인은 같은 클로저로, i를 공유함
nextInt := intSeq()
fmt.Println(nextInt())  // 1
fmt.Println(nextInt())  // 2
fmt.Println(nextInt())  // 3
// 새로운 클로저 생성
newInts := intSeq()
fmt.Println(newInts())  // 1
```

# 44 피보나치 수열

Ex004.go 로 풀었습니다.

# 45 ~ 47 스위치 (Switdch)

일반적인 switch case와 비슷하나, break 가 필요없다.

break 를 하고싶지 않을 때엔 fallthrough 를 사용한다.
```go
switch c := 'c'; c {
case 'c':
    fmt.Println("c is c")
    fallthrough
case 'd':
    fmt.Println("c is c or d")
}
```

case 문은 위에서부터 참인것까지 실행한다.
```go
// Ex)
switch i {
case 0:
case f():
}
-> i == 0 이면 f()는 실행되지 않는다
```

조건을 생략한 switch 는 switch true 와 같다.

-> if-then-else 작성시 깔끔하다.
```go
func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
```

# 48 심화 연습: 복소수 세제곱근

Ex005.go 로 풀었습니다.

# 49 ~ 52 메소드 (Methods)

Go 에서 클래스 대신 struct 에 부착하는 함수를 메소드라 함

-> 함수와 메소드는 다르다!
```go
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
위와 같은 문법으로 메소드를 부착하며, func 키워드와 메소드 이름 사이에

메소드 리시버(method receiver) 를 명시해 준다.

```go
func main() {
    v := &Vertex{3, 4}
    fmt.Println(v.Abs())
}
```
위와 같은 문법으로 메소드를 호출한다.

메소드는 struct 외에 type 에도 붙일 수 있습니다.
```go
type MyFloat float64
func (f MyFloat) Abs() float64 {
   if f<0 {
       return float64(-f)
   }
   return float64(f)
}
```

포인터 리시버 (v *Vertex) 와 값 리시버 (f Myfloat)의 차이는

기억 공간을 그대로 사용하느냐 vs 값을 복사하여 사용하느냐 이다.

값을 복사하는 경우 복사하는 오버헤드를 감수해야하고,

원래 공간의 값을 수정할 수 없습니다.
