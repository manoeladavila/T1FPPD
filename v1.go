// Os filósofos podem entrar em deadlock

package main

import (
     "fmt"
     "sync" // *
     "time"
)

var forks [5]sync.Mutex // * um mutex por garfo

func Philosopher(id, iteration int) {
     time_to_think := 2
     time_to_eat := 3

     // think
     for s := 1; s <= id*6; s++ {
          fmt.Printf(" ")
          }
     fmt.Printf(" T%d\n", iteration) // Thinking
     time.Sleep(time.Duration(time_to_think) * time.Second)

     // * pega o garfo esquerdo e depois o direito
     left := id
     right := (id + 1) % 5
     forks[left].Lock()
     time.Sleep(100 * time.Millisecond) // * garante que o deadlock aconteça, espera circular
     forks[right].Lock()

     // eat

     for s := 1; s <= id*6; s++ {
          fmt.Printf(" ")
          }
     fmt.Printf(" E%d\n", iteration) // Eating// V1 - Versão concorrente sujeita a deadlock


func main() {
	rounds := 5

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go Philosopher(i, rounds, &wg)
	}

	wg.Wait() // espera todos terminarem
	elapsed := time.Since(start)

	fmt.Printf("\nConcurrent Dinner took %s\n\n", elapsed)
}
     time.Sleep(time.Duration(time_to_eat) * time.Second)

     // NOVO: devolve os garfos
     forks[right].Unlock()
     forks[left].Unlock()
}