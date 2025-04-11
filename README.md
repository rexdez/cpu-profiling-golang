# CPU profiling in Golang

This repo is a collection of all the variations I have tried while profiling different processes.
I generates output to folder `pprof-output`. You can run the following command to visualize CPU usage:

For CLI:
```bash
go tool pprof "output/{profile_variation_name}/{profiling_type}.pprof"
```

For http _(ensure graphwiz is installed)_:
```bash
go tool pprof -http=":8000" "output/{profile_variation_name}/{profiling_type}.pprof"
```
replace `filename` with the filename you want to monitor.


## Variations
1. [`for{}` bs `select{}` profiling](https://github.com/rexdez/cpu-profiling-golang/blob/main/variations/for_vs_select.go)
2. [Worker Pool profiling](https://github.com/rexdez/cpu-profiling-golang/blob/main/variations/workerPool.go)
2. [Swinging Goroutine with channels](https://github.com/rexdez/cpu-profiling-golang/blob/main/variations/swinging_goroutine_block.go)


## Results
### 1. `for{}` bs `select{}` profiling
- `select{}` produces no output on CPU usage
- `for{}` process produces 95-100% of CPU usage
