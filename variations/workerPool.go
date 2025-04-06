package variations

import (
	"container/heap"
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime"
	"sync"
	"time"
)

const JOB_LENGTH = 10000

type Job struct {
	val int
	Priority int
	Ctx context.Context
	CancelCtx context.CancelFunc
	Ack bool
}

type JobQueue struct {
	Jobs []*Job
}

func (jq JobQueue) Len() int{
	return len(jq.Jobs)
}

func (jq JobQueue) Less(i, j int) bool {
	return jq.Jobs[i].Priority < jq.Jobs[j].Priority
}

func (jq *JobQueue) Pop() any {
	job := jq.Jobs[len(jq.Jobs)-1]
	jq.Jobs = jq.Jobs[:len(jq.Jobs)-1]
	return job
}

func (jq *JobQueue) Push(v any) {
	val, _ := v.(*Job)
	jq.Jobs = append(jq.Jobs, val)
}

func (jq *JobQueue) Swap(i, j int) {
	jq.Jobs[i], jq.Jobs[j] = jq.Jobs[j], jq.Jobs[i]
}


func ProfileWorkerPool() {
	queue := &JobQueue{
		Jobs: make([]*Job, 0),
	}
	// Using heap to manage jobs based on priority
	heap.Init(queue)
	// in case we need to arbitrarily cancel a job
	var CancelCtxs []context.CancelFunc

	// Creating Jobs and pushing into JobQueue struct
	for range JOB_LENGTH {
		ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
		CancelCtxs = append(CancelCtxs, cancelCtx)
		heap.Push(queue, &Job{
			val: rand.Intn(1000),
			Priority: rand.Intn(3),
			Ctx: ctx,
			CancelCtx: cancelCtx,
			Ack: false,
		})
	}

	// Creating profiling file
	f1, err := os.OpenFile("output/worker-pattern/cpu.pprof", os.O_CREATE | os.O_RDWR | os.O_TRUNC, 0666)
	if err !=nil {
		panic(err)
	}
	defer f1.Close()

	// Start Porfiling Jobs
	pprof.StartCPUProfile(f1)
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	// Initiating Job Receiver
	JobChan := make(chan *Job, JOB_LENGTH)

	// Waitgroup to wait for jobs to complete
	var wg sync.WaitGroup

	// Intializing the worker pool
	for range 10{
		wg.Add(1)
		go Worker(JobChan, &wg)
	}

	var cancelCount int
	// Sending jobs to workers using JobChan
	go func() {
		for i := range JOB_LENGTH {
			val :=  heap.Pop(queue)
			JobChan <- val.(*Job)

			// Randomly cancelling an job
			if getRandomChance() {
				cancelCount++
				CancelCtxs[i]()
			}
		}
		close(JobChan)
	}()

	wg.Wait()
	fmt.Printf("\nTotal jobs cancelled: %v", cancelCount)
	pprof.StopCPUProfile()
	SaveProfile("output/worker-pattern/heap.pprof", "heap")
	SaveProfile("output/worker-pattern/goroutine.pprof", "goroutine")
	SaveProfile("output/worker-pattern/block.pprof", "block")
	SaveProfile("output/worker-pattern/mutex.pprof", "mutex")
	SaveProfile("output/worker-pattern/allocs.pprof", "allocs")
}

func Worker(JobChan chan *Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range JobChan {
		if job.Ctx.Err() != nil {
			fmt.Printf("\nJob Cancelled: %v", job.val)
			continue
		}
		job.val = logRunningTask(job.val)
		job.Ack = true
	}
}

func getRandomChance() bool {
	return rand.Intn(100) > 95
}

func logRunningTask(n int) int {
	var result float64
	for i := range n {
		result += math.Sqrt(float64(n*i+1))
	}
	return int(result)
}