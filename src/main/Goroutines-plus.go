// Goroutines-plus
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numberChan1 := make(chan int64, 3) // 数字通道1。
	numberChan2 := make(chan int64, 3) // 数字通道2。
	numberChan3 := make(chan int64, 3) // 数字通道3。

	// 用于等待一组操作执行完毕的同步工具。
	var waitGroup sync.WaitGroup
	// 该组操作的数量是3。
	waitGroup.Add(3)

	go func() { // 数字过滤装置1。
		for n := range numberChan1 { // 不断的从数字通道1中接收数字，直到该通道关闭。
			if n%2 == 0 { // 仅当数字可以被2整除，才将其发送到数字通道2.
				numberChan2 <- n
			} else {
				fmt.Printf("Filter %d. [filter 1]\n", n)
			}
		}
		close(numberChan2) // 关闭数字通道2。
		waitGroup.Done()   // 表示一个操作完成。
	}()

	go func() { // 数字过滤装置2。
		for n := range numberChan2 { // 不断的从数字通道2中接收数字，直到该通道关闭。
			if n%5 == 0 { // 仅当数字可以被5整除，才将其发送到数字通道2.
				numberChan3 <- n
			} else {
				fmt.Printf("Filter %d. [filter 2]\n", n)
			}
		}
		close(numberChan3) // 关闭数字通道3。
		waitGroup.Done()   // 表示一个操作完成。
	}()

	go func() { // 数字输出装置。
		for n := range numberChan3 { // 不断的从数字通道3中接收数字，直到该通道关闭。
			fmt.Println(n) // 打印数字。
		}
		waitGroup.Done() // 表示一个操作完成。
	}()

	for i := 0; i < 30; i++ { // 先后向数字通道1传送100个范围在[0,30)的随机数。
		numberChan1 <- rand.Int63n(30)
	}
	close(numberChan1) // 数字发送完毕，关闭数字通道1。

	waitGroup.Wait() // 等待前面那组操作（共3个）的完成。
}
