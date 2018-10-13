package main

import "fmt"

//■関数型
//関数を変数に代入するには、関数型を使う

func main() {

    var f func(int, int) int

    f = func(a int, b int) int {
        return a + b
    }

    fmt.Println(f(1, 2))

    //関数型の変数に別の関数を代入
    f = multiply

    fmt.Println(f(1, 2))
}

func multiply (a int, b int) int {
    return a * b
}