package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func main() {
	try_queue()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func count_with_channel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	
	close(c)
}

func try_two_goroutines() {
fmt.Println(runtime.GOOS)
	go count("sheep")
	go count("fish")
	time.Sleep(time.Second * 2)
	//fmt.Scanln()
}

// v mainu, kde je jedina gorutina, pockej az ta potvora dobehne
func try_waitgroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("kangaroo")
		wg.Done()
	} ()
	
	wg.Wait()
}

// neprinti v countu, ale posli msg prec channel
func try_channel() {
	c := make(chan string)
	go count_with_channel("butterfly", c)
	
	for {	
		msg, open := <- c

		if !open {
			break
		}

		fmt.Println(msg)
	}
}

// syntactic sugar for !open check
func try_channel_sugar() {
	c := make(chan string)
	go count_with_channel("butterfly", c)
	
	for msg := range(c) {
		fmt.Println(msg)
	}
}

// deadlock, pac mame jedinou gorutinu ktera se zasekne na
// tom odesilani "hello"
func try_more_channel() {
	c := make(chan string)
	c <- "hello"

	msg := <- c
	fmt.Println(msg)
}

// buffered channel - neblokuje, dokud se nenaplni
// tzn tady neblokuj pr max 2 zpravy
func try_buffered_channel() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "cype"

	msg := <- c
	fmt.Println(msg)

	// tady bez dvojtecky
	msg = <- c
	fmt.Println(msg)
}

// vyber receivnuti gorutiny, ktera sendnula - nesekvencne
func try_select_channel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {	
		for {
			c2 <- "every 2000ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

// kdyby tu nebyl select, ale jenom dva printy, tak by
// bezely sekvencne - tzn ten rychlejsi by cekal na ukor
// toho pomalejsiho
	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}
}


// 
func try_queue() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)

	for i:= 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(j, <-results)
	}
}

// jobs - jenom receive channel
// results - jenom send channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}
