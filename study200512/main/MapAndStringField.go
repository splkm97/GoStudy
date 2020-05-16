package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    // 맵의 생성
    m := make(map[string]int)
    
    // s를 공백을 기준으로 나눔
    arr := strings.Fields(s)
    // s의 각 단어 개수만큼 반복
    for i := 0; i < len(arr); i ++ {
        // arr에 있는 단어를 키값으로하여 밸류값(단어가 나온 개수)을 하나 증가시킴
        m[arr[i]] += 1
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}
