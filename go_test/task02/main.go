package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ✅指针
// 题目 1：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func ques01(p *int) {
	*p = 123
}

func test1() {
	var pr *int
	i := 100
	pr = &i
	fmt.Println(*pr)
	ques01(pr)
	fmt.Println(*pr)
}
func test11() {
	i := 100
	fmt.Println(i)
	ques01(&i)
	fmt.Println(i)
}

// 题目 2：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func ques02(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] = (*p)[i] * 2
	}

}
func test2() {
	var pr *[]int
	arr := []int{1, 2, 3, 4, 5}
	pr = &arr
	fmt.Println(*pr)
	ques02(pr)
	fmt.Println(*pr)
}

func test22() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	ques02(&arr)
	fmt.Println(arr)
}

// ✅Channel
// 题目 3：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

func ques03() {
	ch := make(chan int)

	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 两个协程

	// 协程1：生成1到10的整数并发送到通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 协程2：从通道接收整数并打印
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
	}()

	// 等待两个协程都完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}

// 题目 4：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func ques04() {
	ch := make(chan int, 10)

	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 两个协程

	// 协程1：生成1到10的整数并发送到通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 协程2：从通道接收整数并打印
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
	}()

	// 等待两个协程都完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}

// ✅锁机制
// 题目 5：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

var count int = 1

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取计数
func (c *SafeCounter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func ques05() {
	counter := SafeCounter{}
	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(10) // 10个协程
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println("结束运行 c=", counter.getCount())
}

// 题目 6：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func ques06() {
	var counter int64 = 0
	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(10) // 10个协程
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()

	fmt.Println("结束运行 c=", counter)
}

func main() {
	//my
	// test1()
	// test2()

	//other
	// test11()
	// test22()

	// ques03()
	// ques04()
	// ques05()
	ques06()
}
