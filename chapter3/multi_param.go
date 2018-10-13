package main

import "fmt"

func main() {
    f2(f1())
}

func f1() (int, string, float32){
    return 0, "xyz", 3.14
}

//f1の戻り値と同じパラメータをもつ関数
//パラメータcはfloat32とは異なるが、interface型なので、どんな値でも代入できる
func f2(a int, b string, c interface{}) {
    fmt.Println(a, b, c)
}