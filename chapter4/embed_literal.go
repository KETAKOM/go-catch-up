package main

import "fmt"

//■構造体リテラルにおける埋め込み
//構造体の埋め込みを使用した場合も、宣言と同様に構造体リテラルを使って初期化することができる

type Person struct {
    name string
    age int
}

type Employee struct {
    id int
    Person
}

func main() {
    //構造体リテラルで初期化
    e := Employee{1, Person{name: "太郎", age: 20}}

    fmt.Println(e)
}