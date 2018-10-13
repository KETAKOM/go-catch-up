package main

import "fmt"

//■メソッドのレシーバをポインタにする
//メソッドではレシーバ変数を使って、メソッドを呼び出した対象の値にアクセスすることができる


type myType int

//レシーバが値(非ポインタ)のメソッド
func (value myType) setByValue(newValue myType) {
    value = newValue
}

func (value *myType) setByPointer(newValue myType) {
    *value = newValue
}

func main() {
    var x myType = 0

    //このメソッドのレシーバは値なので、値の変更ができない
    x.setByValue(1)
    fmt.Println(x)

    //このメソッドのレシーバはポインタなので、値の変更ができる
    x.setByPointer(2)
    fmt.Println(x)
}


