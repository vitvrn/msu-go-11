package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Add(x int64) {
	atomic.AddInt64(&c.val, x)

}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounter{}

	// var wg sync.WaitGroup
	// wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
				//runtime.Gosched()
			}
			//wg.Done()
		}(i)
	}

	time.Sleep(time.Second)
	// wg.Wait()
	fmt.Println(counter.Value())
}
