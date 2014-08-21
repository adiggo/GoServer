package main

import "fmt"
//if x, y, z type same, then just declare once
func add(x, y, z int) int{
    return x + y + z
}

func main(){
    fmt.Println(add (40, 13, 23))
}
