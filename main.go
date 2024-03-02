package main

import (
	"fmt"
	"sync"
)

func main() {

	// errorHandlingInIndividualWaitGroup([]int{1, 3, 8})
	errorHandlingWithErrorChannel([]int{1, 3, 8})

}

func errorHandlingInIndividualWaitGroup(indexesWithError []int) {
	WAITGROUP_SIZE := 10

	var wg sync.WaitGroup

	for i := 0; i < WAITGROUP_SIZE; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()

			for _, index := range indexesWithError {
				if i == index {
					fmt.Println("ERROR FROM WITHIN WAIT GROUP: ", i)
					return
				}
			}

			fmt.Println("Hello from waitgroup", i)

		}(i)
	}

	wg.Wait()
}

func errorHandlingWithErrorChannel(indexesWithError []int) {

	WAITGROUP_SIZE := 10

	var wg sync.WaitGroup
	var errorChannel = make(chan error, WAITGROUP_SIZE)

	for i := 0; i < WAITGROUP_SIZE; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()

			for _, index := range indexesWithError {
				if i == index {
					errorChannel <- fmt.Errorf("CHANNEL ERROR: %d", i)
					return
				}
			}

			fmt.Println("Hello from channel waitgroup", i)

		}(i)
	}

	wg.Wait()

	close(errorChannel)

	for err := range errorChannel {
		fmt.Println(err)
	}

}
