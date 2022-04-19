package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const concurrencyProcesses = 10 // limit the maximum number of concurrent reading process tasks
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)
	found := make(chan int)
	limitCh := make(chan struct{}, concurrencyProcesses)

	for i := 0; i < jobCount; i++ {
		go func() {
			limitCh <- struct{}{}
		}()
		//limitCh <- struct{}{} // 卡住了主流程
		go func(val int) {
			//defer func(val int) {
			//	fmt.Println("release channel of ", val)
			//	<-limitCh
			//	//wg.Done()
			//}(val)
			waitTime := rand.Int31n(1000)
			fmt.Println("job:", val, "wait time:", waitTime, "millisecond")
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			found <- val // 在这个channel被接收之后才执行defer
			<-limitCh
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(found)
	}()
	var results []int
	for p := range found {
		fmt.Println("Finished job:", p)
		results = append(results, p)
	}
	//close(found)
	fmt.Println("result:", results)
}
