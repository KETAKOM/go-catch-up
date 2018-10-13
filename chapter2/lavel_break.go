package main

import "fmt"

func main() {

LBL:
    //ループ
    for i := 0; i < 5; i++ {
        switch {

        case i == 3:
            fmt.Println("aaaa", i)

            break LBL


        default:
            fmt.Println(i)
        }
    }

fmt.Println("END")
}