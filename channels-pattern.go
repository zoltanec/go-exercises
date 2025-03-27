package main

import "fmt"
import "sync"

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		defer func() {
			close(channel1)
			close(channel2)
			close(channel3)
		}()

		for j := 0; j < 100; j += 3 {
			channel1 <- j
			channel2 <- j + 1
			channel3 <- j + 2
		}
	}()

	for t := range MergeChannels(channel1, channel2, channel3) {
		fmt.Println(t)
	}
}

func MergeChannels(channels ...<-chan T) <-chan T {
	outputCh := make(chan int)
	var wg = sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				outputCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}
