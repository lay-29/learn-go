package No6

import (
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

func GoroutineTest() {
	go say("world!")
	fmt.Println("Hello")
	time.Sleep(1000 * time.Millisecond)

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all worker done")
}
