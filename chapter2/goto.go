package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 8 {
            goto EVEN
        }
        fmt.Println(i)
    }

    EVEN:
}