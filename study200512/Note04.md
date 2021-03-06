# 57 웹 서버

http 패키지는 http.Handler 를 구현한 어떤 값을 사용해 HTTP 요청을 제공한다.

```go
package http

type Hello struct{}

// http.Handler 구현
func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// http.ListenAndServe 함수로 웹 서버 구동
func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}
```

# 58 연습: HTTP 핸들러

Ex007.go 로 풀었습니다.

# 59 이미지

image 패키지는 Image 인터페이스를 정의합니다.

```go
package image

type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
```

# 60 연습: 이미지

Ex008 로 풀었습니다.

# 61 연습: Rot13 Reader

Ex009 로 풀었습니다.

# 63 고루틴 (Goroutines)

경량 쓰레드로 쓰레드보다 가볍고 사용이 간편함

go 함수명(파라미터) 로 수행합니다.

# 64 채널 (Channels)

채널은 채널 연산자 <- 를 이용해 값을 주고받을 수 있는 파이프이다.

ch := make(chan int) 로 채널 생성

ch <- v 로 채널에 저장

v := <- ch 로 채널에서 꺼냄


이러한 특성이 고루틴에서 락/조건 없이도 동기화 될 수 있도록 도움

# 65 버퍼링되는 채널

채널은 버퍼링될 수 있다. 채널 생성시 make의 인자 두번째에 버퍼 용량을 넣어

버퍼링할 수 있다. 버퍼링하지 않으면 하나의 공간만을 잡는 채널이 생성된다.

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox566415175/prog.go:9 +0x8d
```
버퍼링 범위를 넘어서면 위와 같은 에러가 나온다.

# 66 Range 와 Close

```go
v, ok := <- ch
```
위와 같이 ok를 리턴값을 받으면 채널이 닫혔는지 확인 가능하다

close로 채널을 닫으면 더이상 보낼 값이 없다는것을 전달할 수 있다.

** 송신측만이 채널을 닫을 수 있습니다.

** 채널은 파일과 달라 항상 닫을 필요가 없다.


# 67, 68 셀렉트 (Select)

select 구문은 고루틴이 다수의 통신 동작으로부터 수행 준비를 기다릴 수 있게 한다고 한다.


case 구문으로 받는 동작중 하나가 수행될 수 있을 때까지 수행을 블록합니다.(wait)

다수의 채널이 동시에 준비되면 그 중 하나를 무작위 선택합니다.


select 의 default 케이스는 현재 수행 준비가 완료된 케이스가 없을때 수행됩니다.

블로킹 없이 송/수신하려면 default 케이스를 사용하세요.


# 69, 70 연습: 동등한 이진 트리

Ex010 으로 풀었습니다.
