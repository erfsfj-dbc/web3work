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

const pre int = 1
const a int = iota
const (
	b = 3
	c
	d int = iota
	e
)
const (
	f = 2
	g = iota
	h
	i
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
