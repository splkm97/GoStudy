package main

import "fmt"

func main() {
    for i := 2; i < 10; i ++ {
        fmt.Println("< ", i, "단 >")
        for j := 1; j < 10; j ++ {
            fmt.Printf("%d x %d = %d\n", i, j, i*j)
        }
        fmt.Println()
    }
}
