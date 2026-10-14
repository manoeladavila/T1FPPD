//V3

package main

import (
	"fmt"
	"sync"
	"time"
)

var forks [5]sync.Mutex
var waitTotal [5]time.Duration
var maxWait [5]time.Duration
var eatStart, eatEnd [5][]time.Time
var cycles [5]int

// maior número de refeições que estavam em andamento ao mesmo tempo
func maxEatingSimultaneously() int {
	best := 0
	for i := range eatStart {
		for _, s := range eatStart[i] {
			count := 0
			for j := range eatStart {
				for k := range eatStart[j] {
					if !eatStart[j][k].After(s) && eatEnd[j][k].After(s) {
						count++
					}
				}
			}
			if count > best {
				best = count
			}
		}
	}
	return best
}

func Philosopher(id, iteration int) {
	var left, right int
	time_to_think := 2
	time_to_eat := 3

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)

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

	start := time.Now()
	forks[left].Lock()
	forks[right].Lock()
	end := time.Since(start)
	waitTotal[id] += end
	if end > maxWait[id] {
		maxWait[id] = end
	}
	eatStart[id] = append(eatStart[id], time.Now())

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration)
	time.Sleep(time.Duration(time_to_eat) * time.Millisecond)

	eatEnd[id] = append(eatEnd[id], time.Now())

	forks[right].Unlock() //* devolve o garfo diretiro e dps o esquerdo
	forks[left].Unlock()

	cycles[id]++
}

func main() {
	philosophers := 5
	rounds := 20

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

	var total, maxW time.Duration
	for i := 0; i < philosophers; i++ {
		total += waitTotal[i]
		if maxWait[i] > maxW {
			maxW = maxWait[i]
		}
	}
	fmt.Printf("\nEspera média pelos garfos: %s\n", total/time.Duration(philosophers*rounds))
	fmt.Printf("Maior espera: %s\n", maxW)
	fmt.Printf("Máx. comendo simultaneamente: %d\n", maxEatingSimultaneously())
	fmt.Printf("Ciclos concluídos: F0=%d F1=%d F2=%d F3=%d F4=%d\n", cycles[0], cycles[1], cycles[2], cycles[3], cycles[4])

	fmt.Printf("\nConcurrent Dinner took %s\n\n", elapsed)
}