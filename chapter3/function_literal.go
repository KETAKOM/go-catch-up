package main

import "fmt"

//■関数リテラル

func main() {
    val := 1234

    //関数リテラルの記述と呼び出しを同時に行う
    func (i int) {
        fmt.Println(i * val)
    }(10)

    //関数リテラルを変数に代入
    f := func (s string) {
        fmt.Println(s)
    }

    f("hoge")
}