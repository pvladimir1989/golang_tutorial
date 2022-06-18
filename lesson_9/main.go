package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go say("hello")
// 	fmt.Println("1")
// 	fmt.Println("2")
// 	fmt.Println("3")
// 	fmt.Println("4")
// 	fmt.Println("5")
// 	fmt.Println("6")
// 	time.Sleep(2 * time.Second)
// }

// func say(word string) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(word)
// }

// func main() {
// 	ch := make(chan string)
// 	go say("hello", ch)
// 	fmt.Println("1")
// 	fmt.Println("2")
// 	fmt.Println("3")
// 	fmt.Println("4")
// 	fmt.Println("5")
// 	fmt.Println(<-ch) //  читает exit

// }

// func say(word string, ch chan string) {
// 	time.Sleep(10 * time.Second)
// 	fmt.Println(word)
// 	ch <- "exit"
// }

// func main() {
// 	ch := make(chan int)
// 	go sayHello(ch)
// 	// s := <-ch
// 	fmt.Println(<-ch) //  читает exit
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

// func say(word string) {
// 	fmt.Println(word)

// }

// func sayHello(exit chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		say("hello")
// 		exit <- 1
// 	}

// }

// func main() {
// 	ch := make(chan int)
// 	go sayHello(ch)

// 	for i := range ch {
// 		fmt.Println(i)

// 	}

// }

// func sayHello(exit chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		exit <- 1

// 	}
// 	close(exit)

// }

func main() {
	data := make(chan int)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)

}

func selectOne(data, exit chan int) {
	x := 0

	for {
		select {

		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(50 * time.Millisecond)
		}

	}
}
