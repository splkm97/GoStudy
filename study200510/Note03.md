# 13 변수의 짧은 선언

다음 변수 선언을

```go
var x, y, z int = 1, 2, 3
```

다음과 같이 변수 선언을 할 수도 있다.

```go
x, y, z := 1, 2, 3
```

# 14 상수 (Constants)

const 키워드로 문자, 문자열, 부울, 숫자 상수를 선언 가능하다.

문자 하나를 출력할때에는 fmt.Printf() 함수의 형식지정자 "%c"를 이용해야 하는듯

-> 왜 fmt.Println()에서는 문자 하나 출력시 아스키 코드값을 내보낼까? 질문**

# 15 숫자형 상수 (Numeric Constants)

숫자형 상수는 정밀한 값(int 범위를 벗어나는 값 등)을

** but, fmt.Println()에서는 int범위 벗어나는 값을 출력 불가능

# 16~18 For 반복문

C 나 Java 에서처럼 사용하되, ()가 없다.
```go
for i := 0; i < 10; i ++ {
  /* do something */
}
```
조건문만으로 for 반복을 표현하는것도 가능

(c/Java 에서 while 과 같다)
```go
sum := 1
for sum < 1000 {
  sum += sum
}
```
조건문 생략으로 무한루프 표현
```go
for {
}
```

# 20 조건문 (If)

for 문과 비슷하게 ()없이 사용
```go
if x < 0 {
    fmt.Println("x is smaller than 0")
}
```

# 21, 22 if 와 짧은 명령 사용

if와 조건 사이에 짧은 명령문 하나를 삽입 가능하다.
```go
if v := math.Pow(x, n); v < lim {
// -------------------- 이부분
    fmt.Println(v)
    return v
}
return lim
```

# 23 연습: 루프와 함수

main 디렉 안 ex001.go 로 문제를 풀었음.
