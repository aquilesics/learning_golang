package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	contador := 0
	total_gr := 60

	var wg sync.WaitGroup

	wg.Add(total_gr)

	for i := 0; i < total_gr; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(contador)

}
