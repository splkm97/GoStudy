package main

import (
    "io"
    "os"
    "strings"
)

/*
rot13이란
모든 알파벳의 아스키 값에 13을 더하고
z or Z 를 넘어가면
a or A 로 로테이션한 뒤 더하는 방식의 암호화를 말한다.
*/

// rot13 을 위한 구조체 선언
/*
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/
type rot13Reader struct {
    
    r io.Reader
}

// 13을 각 아스키 값에 더하고 로테이션 적용한다.
func rot13byte(by byte) byte {
    s := rune(by)
    if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
        by += 13
    }
    if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
        by -= 13
    }
    return by
}

// rot13 을 이용해 각 문자를 해석
func (rot13 rot13Reader) Read(data []byte) (len int, err error) {
    len, err = rot13.r.Read(data)
    for i := 0; i < len; i ++ {
        data[i] = rot13byte(data[i])
    }
    
    return 
}


func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
