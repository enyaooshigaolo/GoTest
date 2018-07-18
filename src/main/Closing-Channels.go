// Closing-Channels
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//the more value will be false if jobs has been closed
	//and all values in the channel have already been received.
	//协程不保证执行顺序
	//简单来说，你要等协程完了主进程才要结束掉。所以要等done里灌入数据。
	//你这两个结果，第二次是主协程先退出了,所以整个程序都退出了
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

}
