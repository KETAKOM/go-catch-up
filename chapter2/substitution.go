package main

import "fmt"

func test() (int, int){
    return 10, 20
}

func main() {

    type myString string

    var a, b, c int

    var (
        d       int
        e, f    string
        g, h    myString
    )

    a, b, c, d = 1, 2, 3, 4
    e, f, g, h = "test", "aaaa", "iiii", "uuuu"

    fmt.Println(a, b, c, d)
    fmt.Println(e, f, g, h)



    var aa, bb int
    aa,bb = test()

    fmt.Println(aa, bb)
}