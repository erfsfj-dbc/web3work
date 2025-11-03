// 包声明
package main

// 引入包声明
import "fmt"

// 函数声明
func printInConsole(s string) {
    fmt.Println(s)
}

// 全局变量声明
var str string = "Hello world!"


func main() {
    printInConsole(str)
}