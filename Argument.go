package main

import "fmt"
// from the function, we can see that int is declared after x, int is declared after function add.
func add (x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(43, 13))
}
