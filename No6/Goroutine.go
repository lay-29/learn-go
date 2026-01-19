package No6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}
func worker(goId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker ", goId, " start!")
	time.Sleep(time.Second * 5)
	fmt.Println("worker ", goId, " end!")
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到取消信号，协程退出")
			return // 主动退出
		default:
			fmt.Println("正在工作...")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func GoroutineTest() {
	//1.go
	go say("world!")
	fmt.Println("Hello")
	time.Sleep(1000 * time.Millisecond)

	//2.WaitGroup
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all worker done")

	//3.Context
	ctx, cancel := context.WithCancel(context.Background())

	go doWork(ctx)
	time.Sleep(2000 * time.Millisecond)

	cancel()
	time.Sleep(1 * time.Second)

}
