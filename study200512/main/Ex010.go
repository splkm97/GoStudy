package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    ch <- t.Value
    if t.Left == nil && t.Right == nil {
        return
    }
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    var ch1, ch2 chan int
    var arr [10] int
    cnt := 0
    ch1 = make(chan int, 10)
    ch2 = make(chan int, 10)
    Walk(t1, ch1)
    Walk(t2, ch2)
    for i := 0; i < 10; i ++ {
        arr[i] = <- ch1
    }
    for i:= 0; i < 10; i ++ {
        x := <- ch2
        for j:= 0; j < 10 ; j ++ {
            if x == arr[j] {
                cnt++
            }
        }
    }
    if cnt == 10 {
        return true
    } else {
        return false
    }
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
    
}
