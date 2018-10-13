package main

import "fmt"

//■初期化処理(init関数)

func init() {
    fmt.Println("初期化処理")
}

//一つのファイルに複数のinit関数をもつことができる
func init() {
    fmt.Println("初期化処理2")
}

func main() {
    fmt.Println("main")
}

