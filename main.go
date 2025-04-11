package main

import (
	"fmt"
	"os"
	"profilingGo/variations"
)



func main() {
restart_Choice:
	var user_choice int

	fmt.Printf(`
=======================================
Welcome to the profiler, kindly choose any profiler to run:

-1. Exit()
 1. for{} vs. select{} profiling,
 2. Worker Pool Profiling
 3. Profiling Swinging control flow

Enter choice:`)	
	fmt.Scan(&user_choice)
	switch user_choice {
	case 1:
		variations.ProfileForVsSelect()
	case 2:
		variations.ProfileWorkerPool()
	case 3:
		variations.OddEvenBlock()
	case -1:
		os.Exit(0)
	default:
		fmt.Println("<<<<<<< Invalid Choice >>>>>>>")
		goto restart_Choice
	}
}