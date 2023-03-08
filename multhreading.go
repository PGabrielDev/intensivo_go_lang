package main

import (
	"fmt"
	"time"
)

// Exemplos tasks
// func task(name string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(name, ":", i)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	// go task("tarefa 1 ")
	// go task("tarefa 2 ")
	//  task("tarefa 3 ")
	// canal := make(chan int)

	// go publish(canal)
	// go reader(canal)
	// time.Sleep(time.Second * 10)
	// close(canal)
	ch := make(chan int)

	workers := 3

	//Inicializa os workers
	for i := 1; i <= workers; i++ {
		go worker(i, ch)
	}

	//Joga a carga para os woorkes
	for i := 0; i < 10000; i++ {
		ch <- i
	}
}

//exemplo publicador de menssagem em um canal
// func publish(ch chan int) {
// 	for i := 1; i < 10; i++ {
// 		ch <- i
// 	}
// }

// func reader(ch chan int) {
// 	for x := range ch {
// 		fmt.Println(x)
// 	}
// }

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
