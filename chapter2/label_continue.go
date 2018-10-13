package main

import "fmt"

func main() {

LBL:
    for i := 0; i < 5; i++ {
        fmt.Println(i)

        for j := 0; j < 6; j++ {
            fmt.Println("    ", j)

            if i == 1 && j == 1 {
                continue LBL
            }
        }
    }

fmt.Println("END")
}