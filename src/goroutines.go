package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("cpu cores:%v | n goroutines:%v\n",runtime.NumCPU(), runtime.NumGoroutine())

	wg.Add(2)
	

	go func_1()
	go func_2()


	fmt.Printf("cpu cores:%v | n goroutines:%v\n",runtime.NumCPU(), runtime.NumGoroutine())

	wg.Wait()

}

func func_1() {
	for i := 0; i < 10; i++ {

		fmt.Println("funcao 1:", i)
		time.Sleep(20)

	}
	wg.Done()
}

func func_2() {
	for i := 0; i < 10; i++ {

		fmt.Println("funcao 2:", i)
		time.Sleep(20)

	}
	wg.Done()
}
