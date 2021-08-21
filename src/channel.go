package main

import "fmt"

func main() {
	canal_2 := make(chan int)
	canal := make(chan int)

	go func() {
		canal <- 41

	}()

	fmt.Println(<-canal)

	go meuloop(10, canal_2)

	for v := range canal_2 {
		fmt.Println("canal 2", v)
		// range and close
	}
	// -------------------SELECT EXAMPLE------------------
	x := 500
	a := make(chan int)
	b := make(chan int)

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i

		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i

		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("canal a:", v)
		case v := <-b:
			fmt.Println("canal b", v)
		}
	}
	// -----------------COMA OK-------------------------

	y := make(chan int)

	go func() {
		y <- 42
		close(y)
	}()

	z, ok := <-y

	fmt.Println(z, ok)
	z, ok = <- y

	fmt.Println(z, ok)
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i

	}
	close(s)

}
