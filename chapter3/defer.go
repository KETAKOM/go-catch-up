package main

import "fmt"

//■関数の遅延実行
//defer文は関数の実行を遅延させるときに使う
//deferは呼び出し元の関数の処理が終了してから、実行される

func main() {
    fmt.Println("開始")

    //deployを遅延実行(mainを抜けた後に実行される)
    defer deploy()

    fmt.Println("終了")
}

func deploy() {

    fmt.Println("deplyが呼び出されました。")
}