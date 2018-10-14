package main

import "fmt"

//■構造体

func main() {

    //構造体型の変数を宣言
    var x struct {
        //フィールドの宣言
        i1, i2 int
        f float32
        s string
    }

    //構造体の各フィールドをセット
    x.i1 = 1
    x.i2 = 2
    x.f = 3.14
    x.s = "golang"

    fmt.Println(x)

    //構造体型は便利だが、このままでは使いにくいので、通常は「type」を使用してわかりやすい名前をつける
    type human struct {
        name string
        height int
    }

    var man human
    man.name = "テスト男"
    man.height = 180
    fmt.Println(man)
}