package main

import (
    "net/http"
)

// 에러 부착용 String 타입 지정
type String string

// Struct 구조체 지정
type Struct struct {
    Greeting string
    Punct string
    Who string
}

// String 의 HTTP 요청시 작동하는 메소드
func (s String) ServeHTTP (
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

// Strunct 의 HTTP 요청시 작동하는 메소드
func (s *Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprintf(w, s.Greeting, s.Punct, s.Who)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", stri)
}
