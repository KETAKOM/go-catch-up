package main

import "fmt"

//可変長パラメータ
func main() {
    holiday(1, "元旦", "成人の日")
    holiday(2, "建国記念日")
    holiday(3, "春分の日")
}

//可変長パラメータdaysをもつ関数
//関数の最終パラメータだけは、可変長パラメータとして使用できる
func holiday(month int, days ... string) {
    fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
    for _, day := range days{
        fmt.Println(day)
    }
}