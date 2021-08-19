package main

import (
	"fmt"
	"runtime"
	"sync" 
)

func main() {
	var mu sync.Mutex

	contador := 0
	total_gr := 100

	var wg sync.WaitGroup

	wg.Add(total_gr)

	for i := 0; i < total_gr; i++ {
		go func() {

			mu.Lock() //exclusao mutua
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock() // excluasao multipla

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(contador)

}
