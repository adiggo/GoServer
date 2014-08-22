package main

import "fmt"
//in go language, type can be not declared. 
//it will catch the type based on initilizer
var i, j = 1, 2

func main(){
    var c, python = true, "no";
    fmt.Println(i, j, c, python)
}
