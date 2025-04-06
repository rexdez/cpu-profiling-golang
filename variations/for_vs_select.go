package variations

import (
	"os"
	"runtime/pprof"
	"time"
)

// Profiles the CPU for `select{}` & `for{}` indefinite blocking methods
func ProfileForVsSelect() {
	// Initializing profile files using absolute paths
	f1, err := os.OpenFile("output/for-select/select-pprof.pprof", os.O_CREATE | os.O_RDWR, 0666)
	if err !=nil {
		panic(err)
	}
	f2, err := os.OpenFile("output/for-select/for-pprof.pprof", os.O_CREATE | os.O_RDWR, 0666)
	if err !=nil {
		panic(err)
	}

	ch := make(chan bool)
	// Closing files
	defer f1.Close()
	defer f2.Close()

	// Initiating profiling for `select{}` indefinite block
	pprof.StartCPUProfile(f1)
	go ShowProgress(ch)
	go func() {
		select{}
	}()

	// Waits for 10 second to profile the above method and exiting
	<- time.After(time.Second*10)
	ch <- true
	pprof.StopCPUProfile()
	
	// Repeating the above method for profiling `for{}` blocking method
	pprof.StartCPUProfile(f2)
	go ShowProgress(ch)
	go func() {
		for{}
	}()
	<- time.After(time.Second*10)
	ch <- true
	pprof.StopCPUProfile()
}