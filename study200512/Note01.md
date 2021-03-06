# 24 기본 자료형

```go
bool
string

int  int8  int16  int32  int64  // 정수형
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 다른 이름(alias)

rune // int32의 다른 이름(alias)
     // 유니코드 코드 포인트 값을 표현합니다. 

float32 float64 // 실수형

complex64 complex128  // 허수 표현을 위함.
```

# 25-26 구조체 (Structs)

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

# 27 포인터 (Pointers)

Go 에는 포인터가 있지만, 포인터 연산은 불가능하다

# 28 구조체 리터럴

{} 안에 필드와 값들을 나열하여 새로 할당.

{Name: value} 를 통해 할당도 가능

&을 사용하면 구조체 리터럴 포인터 생성 가능


# 29 new 함수

`new(T)` 는 zero value 가 할당된 T타입 포인터 반환합니다.

```go
var t *T = new(T)
t := new(T)
```

# 30-36 슬라이스 (Slices)

슬라이스는 배열의 값을 가리킨다.

사용법
```go
p = []int{1, 2, 3, 4}
p[2] -> 3
p[1:3] -> {2, 3}
p[:2] -> {1, 2}
p[2:] -> {3, 4}
```
```go
make : len() -> 길이(0으로 초기화) cap() -> 용량(len 보다 큰 용량은 nil)
a := make([]int, 5) // len(a) -> 5
b := make([]int, 0, 5)   // len(b) = 0, cap(b) = 5

b = b[:cap(b)] // len(b) = 5, cap(b) = 5
b = b[1:]      // len(b) = 4, cap(b) = 4
```
range 를 사용해 슬라이스나 맵의 순회가 가능하다.
```go
func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```
range 를 사용한 for 문에서 i, value 를 설정할 때
1) i, value 모두 사용
```go
for i, value := range pow {
     fmt.Printf("pow[%d] = %d\n", i, value)
}
for i := range pow {
     pow[i] = 1 << uint(i)
}
for _, value := range pow {
     fmt.Printf("%d\n", value)
}
```

# 37 맵 (Maps)

맵: Python의 딕셔너리와 비슷

사용법
```go
var m map[string]int
m = make(map[string]int)
m["hello"] = 40
```

# 38,39 맵 리터럴 (Map literals)

구조체 리터럴과 비슷 but key 값을 반드시 지정해야함
```go
var m = map[string]int {
     "hello": 10,
     "bye": 20,
}
```
가장 상위 타입이 타입명이면 리터럴에서 타입명 생략이 가능합니다.
```go
// Ex)
"Bell Labs": {40.68433, -74.39967}
"Bell Labs": Vertex{40.68433, -74.39967}
```
두 표현 모두 같은 표현입니다.

# 40 맵 다루기 (Mutating Map)

맵은 다음과 같은 방법으로 사용한다.

-----------------------------------------------
m[key] = elem

-> 맵 m의 요소 삽입/수정

elem = m[key]

-> 맵 m의 요소 가져오기

delete(m, key)

-> 맵 m의 key요소 삭제

elem, ok = m[key]

-> 맵 m에 key에 해당하는 값 존재하는지 확인

-> 있다면? -> ok = true / elem = m[key]

-> 없다면? -> ok = false / elem = 0 (zero value)

-----------------------------------------------
