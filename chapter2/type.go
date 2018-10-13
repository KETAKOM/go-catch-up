package main

import "fmt"

func main() {

    //int型からmyInteger型を定義
    type myInteger int

    //myInteger型の変数を定義
    var i myInteger = 123;

    i += 1

    fmt.Println(i)

    //新しい構造体を作成
    type myStruct struct {
        a int
        b int
    }
}