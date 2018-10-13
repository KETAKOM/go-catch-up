package main

import "fmt"

func main () {
    //定数を定義
    const G float64 = 3.14159

    //複数の定数を定義
    const a, b, c int = 1 ,2 ,3

    const (
        ja string = "日本語"
        en string = "english"
    )

    const (
        x int = 1
        y
        z
    )

    fmt.Println(G)
    fmt.Println(a, b, c, ja, en, x, y, z)
}