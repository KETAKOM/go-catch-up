package main

import "fmt"

//■メソッドの宣言
//Golangでは型に「メソッド」を持たせることができる

//int型から新たにmyType型を宣言
type myType int

//myTypeをレシーバにもつ関数、つまりmyType型のメソッドを宣言
func (value myType) hogehoge(ii myType) {
    fmt.Println(value + ii)
}


func main() {

   var a myType = 99

   //myType型のメソッドを呼び出し
   a.hogehoge(1)

}