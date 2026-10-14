//V0

// Example code filo_seq.go - Sequential Dinning Philosophers
// Since is sequential, elapsed timne could be simply calculated by (roubnds*philosophers*(time_to_think+time_to_eat))
// By Cesar De Rose

package main

import (
	"fmt"
	"time"
)

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
	time_to_think := 2
	time_to_eat := 3

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)

	eatStart[id] = append(eatStart[id], time.Now())

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration) // Eating
	time.Sleep(time.Duration(time_to_eat) * time.Millisecond)

	eatEnd[id] = append(eatEnd[id], time.Now())

	cycles[id]++
}

func main() {
	philosophers := 5
	rounds := 20

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run philosophers

	start := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < philosophers; j++ {
			Philosopher(j, i)
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("\nEspera média pelos garfos: 0s (sem disputa: execução sequencial)\n")
	fmt.Printf("Maior espera: 0s (sem disputa: execução sequencial)\n")
	fmt.Printf("Máx. comendo simultaneamente: %d\n", maxEatingSimultaneously())
	fmt.Printf("Ciclos concluídos: F0=%d F1=%d F2=%d F3=%d F4=%d\n", cycles[0], cycles[1], cycles[2], cycles[3], cycles[4])

	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}