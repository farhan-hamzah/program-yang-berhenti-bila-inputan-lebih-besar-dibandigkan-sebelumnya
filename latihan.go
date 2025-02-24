package main

import "fmt"

func main() {
    var count, num, last int
    fmt.Scan(&num)
    last = num
    count = 1
    for count < 10 {
        fmt.Scan(&num)
        if num > last {
            break
        }
        last = num
        count++
    }
    fmt.Println(count) 
}
