package main

import "fmt"

//メソッドは通常、メソッドを実装している型の値に対して呼び出すが、
//「値」と「メソッド」の紐付けをメソッド値として扱うことができる

type myType int

func (value *myType) add(inc myType) myType {
    *value += inc
    return *value
}

func main() {

    var i myType

    //変数iに対し、addメソッドを呼び出し、3を加算
    fmt.Println(i.add(3))

    //メソッド値を取得
    mv := i.add

    fmt.Println(mv(3))
}

