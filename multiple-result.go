package main

import "fmt"
// we can return miltiple value , big difference from java
func swap(x, y string) (string, string){
    return y, x
}

func main(){
    a, b := swap ("hello", "world")
    fmt.Println(a, b)
}
