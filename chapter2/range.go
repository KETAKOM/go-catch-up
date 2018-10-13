package main

import "fmt"

func main() {
    //配列の初期化
    arr := [...]int{0, 1, 2, 3}

    for i := range arr {
        fmt.Println(i)
    }
}