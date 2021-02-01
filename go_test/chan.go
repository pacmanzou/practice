package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var catCounter uint64
	var dogCounter uint64
	var fishCounter uint64

	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	wg.Add(3)

	go CatShow(&wg, catChan, dogChan, catCounter)
	go DogShow(&wg, dogChan, fishChan, dogCounter)
	go FishShow(&wg, fishChan, catChan, fishCounter)

	catChan <- struct{}{}
	wg.Wait()
}

func CatShow(wg *sync.WaitGroup, catChan, dogChan chan struct{}, counter uint64) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-catChan
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogChan <- struct{}{}
	}
}

func DogShow(wg *sync.WaitGroup, dogChan, fishChan chan struct{}, counter uint64) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-dogChan
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishChan <- struct{}{}
	}
}

func FishShow(wg *sync.WaitGroup, fishChan, catChan chan struct{}, counter uint64) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-fishChan
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catChan <- struct{}{}
	}
}
