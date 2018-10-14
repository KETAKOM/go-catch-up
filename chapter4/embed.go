package main

import "fmt"

//■構造体における特殊なフィールドについて
//構造体のフィールド名として「ブランク識別子」を使用することができる。
//これはブランクフィールドと呼ばれる
//■匿名フィールドの書き込み
//フェールドの宣言が型だけの時(フィールド名を省略した時)を「匿名フィールド」と呼ぶ
//匿名フィールドは「埋め込み」と呼ばれる


//埋め込まれる側の構造体
type embedded struct {
    i int
}

//embedded型のメソッド
func (x embedded) doSomething() {
    fmt.Println("doSomething")
}

//埋め込み先の構造体
type test struct {
    embedded
}

func main() {
    var x test
    x.doSomething()

    x.i = 100

    fmt.Println(x.i)

}



