package main

import (
	"fmt"
	"sync"
	"time"
)

var forks [5]sync.Mutex

func Philosopher(id, iteration int) {
	var left, right int
	time_to_think := 2
	time_to_eat := 3

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time.Duration(time_to_think) * time.Second)

	// * se o id do filosofo for par, pega o fork left depois o fork right. Se o id for impar, pega o right, depois o left.
	//Isso evita que acontece o deadlock, porque não existe a possibilidade de esperarem eternamente pelo mesmo recurso.
	// Ainda existe a possibilidade de ocorrer starvation.

	if id%2 == 0 {
		left = id
		right = (id + 1) % 5
	} else {
		left = (id + 1) % 5
		right = id
	}

	forks[left].Lock()
	time.Sleep(100 * time.Millisecond) //* não ocorre mais deadlock. Os filósofos nunca vão entrar em espera circular, porque 3 filósogos sempre vão pegar o garfo direito antes do esquerdo.
	forks[right].Lock()

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration)
	time.Sleep(time.Duration(time_to_eat) * time.Second)

	forks[right].Unlock() //* devolve o garfo diretiro e dps o esquerdo
	forks[left].Unlock()
}

func main() {
	philosophers := 5
	rounds := 5

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run philosophers

	start := time.Now()

	var wg sync.WaitGroup
	for j := 0; j < philosophers; j++ { // for faz uma goraoutine por filosofo
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 1; i <= rounds; i++ {
				Philosopher(id, i)
			}
		}(j)
	}
	wg.Wait() // espera todos terminatrem

	elapsed := time.Since(start)

	fmt.Printf("\nConcurrent Dinner took %s\n\n", elapsed)
}
