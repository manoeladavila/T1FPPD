//com sincronização excessiva
//a mesa toda vai ser um mutes (seção crítica) e só um filósofo come por vez
//é diferente da v0 (sequencial) pq os filosofos ainda pensam concorrente
package main

import (
	"fmt"
	"sync"
	"time"
)

// var forks [5]sync.Mutex   (não precisamos de garfos, a mesa toda é um mutex agr)
var table sync.Mutex

func Philosopher(id, iteration int) {
	time_to_think := 2
	time_to_eat := 3

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time.Duration(time_to_think) * time.Second)

	table.Lock()

	/* * pega o garfo esquerdo e dps o direito
	left := id
	right := (id + 1) % 5 // pq o filosofo 4 tem que pegar o garfo 0 (4+1%5)
	forks[left].Lock()
//	time.Sleep(100 * time.Millisecond) //* força que o deadlock aconteça, espara circular, pq todos os filosofos vao pegar o garfo da esquerda primeiro
	forks[right].Lock() */

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration)
	time.Sleep(time.Duration(time_to_eat) * time.Second)

	/*forks[right].Unlock()  //* devolve o garfo diretiro e dps o esquerdo
	forks[left].Unlock() */

	table.Unlock()
}

func main() {
	philosophers := 5
	rounds := 5

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run philosophers

	start := time.Now()

	var wg sync.WaitGroup
	for j := 0; j < philosophers; j++ {  // for faz uma goraoutine por filosofo
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



//interprentando os resultados, cada filósof entra no mutex come e volta a pensar, T2 é sempre imeddiatamente depois de E1
