package main

import (
    "fmt"
//_ 匿名导入, 不能访问 pkg1 中的任何函数
    _ "github.com/learn/init_order/pkg1"
)

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
    fmt.Println("main init method invoked")
}

func main() {
    fmt.Println("main method invoked!")
}

func getMainVar() string {
    fmt.Println("main.getMainVar method invoked!")
    return mainName
}