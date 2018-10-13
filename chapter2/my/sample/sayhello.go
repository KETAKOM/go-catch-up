package sample  // 「my/sample」だが最下層のパッケージ名だけを記述する


import "fmt"

func SayHello(who string) {

    fmt.Println("HELLO " + who)
}