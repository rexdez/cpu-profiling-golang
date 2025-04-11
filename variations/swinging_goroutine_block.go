package variations

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"runtime"
)

func OddEvenBlock() {
	// Initating pprof file
	f, err :=  os.OpenFile("output/goroutine-block/cpu.pprof", os.O_TRUNC | os.O_RDWR | os.O_CREATE, 0666)
	
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	pprof.StartCPUProfile(f)
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	// Channels for swinging control flow
	oddCh := make(chan bool)
	evenCh := make(chan bool)

	limit := 1000
	// waitgroup to wait for go routine to stop
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(oddCh, evenCh, &wg, limit)
	go printEven(oddCh, evenCh, &wg, limit)	
	
	// Starting first swing on the odd side
	oddCh <- true

	wg.Wait()

	pprof.StopCPUProfile()
	SaveProfile("output/goroutine-block/heap.pprof", "heap")
	SaveProfile("output/goroutine-block/goroutine.pprof", "goroutine")
	SaveProfile("output/goroutine-block/block.pprof", "block")
	SaveProfile("output/goroutine-block/mutex.pprof", "mutex")
	SaveProfile("output/goroutine-block/allocs.pprof", "allocs")
}

func printOdd(oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i = i+2 {
		<- oddChan
		if i&1 == 1 {
			fmt.Println(i)
		}
		evenChan <- true
	}
	<- oddChan
}

func printEven(oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 2; i <= n; i=i+2 {
		<- evenChan
		if i&1 == 0 {
			fmt.Println(i)
		}
		oddChan <- true
	}
}