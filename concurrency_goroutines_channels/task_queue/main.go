package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct for Task queue
type Task struct {
	ID      int
	Execute func() error
	Retry   int
}

type WorkerPool struct {
	TaskQueue   chan Task
	TotalWorker int
	wg          sync.WaitGroup
	RetryLimit  int
}

// initialize the worker pool
func newWorkerPool(numWorker, retryLimit int) *WorkerPool {
	return &WorkerPool{
		TaskQueue:   make(chan Task),
		TotalWorker: numWorker,
		RetryLimit:  retryLimit,
	}
}

// method to assign task to each worker
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.TotalWorker; i++ {
		wp.wg.Add(1)    // add the initial worker to wait group
		go wp.worker(i) // calls the go routine for each i
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.TaskQueue {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		err := task.Execute()
		retries := 0
		for err != nil && retries < wp.RetryLimit {
			fmt.Printf("Task %d failed retrying...(%d%d)\n", task.ID, retries+1, wp.RetryLimit)
			retries++
			err = task.Execute()
		}
		if err != nil {
			fmt.Printf("Task %d failed even after %d retries", task.ID, retries)
		} else {
			fmt.Printf("Worker %d completed task %d", id, task.ID)
		}
	}
}

// Add task to the queue
func (wp *WorkerPool) AddTask(task Task) {
	wp.TaskQueue <- task
}

// Graceful shutdown of each worker
func (wp *WorkerPool) Stop() {
	close(wp.TaskQueue)
	wp.wg.Wait()
}

// Dynamic(manual) resizing of the number of workers working
func (wp *WorkerPool) Resize(newSize int) {
	/*  If newSize > TotalWorker then increase the workers by newSize - TotalWorker
	newSize - TotalWorker == naye Majdoor needed,start majdoors func for these	*/
	if newSize > wp.TotalWorker {
		diff := newSize - wp.TotalWorker
		for i := 0; i < diff; i++ {
			wp.wg.Add(1)
			go wp.worker(i)
		}
		wp.TotalWorker = newSize
	} else if newSize < wp.TotalWorker {
		diff := wp.TotalWorker - newSize
		for i := 0; i < diff; i++ {
			/*  signals the majdoors to stop acc to newSize worker start when size === diff they stop cause we
			want this much majdoors only */
			wp.TaskQueue <- Task{ID: -1, Execute: nil}
		}
		wp.TotalWorker = newSize

	}
}

// Automatic scaling based on Queue Load
func (wp *WorkerPool) AutoResize() {
	for {
		queueSize := len(wp.TaskQueue)
		if queueSize > 50 && wp.TotalWorker < 10 {
			fmt.Println("Increasing worker as load has increased")
			wp.Resize(wp.TotalWorker + 1)
		} else if queueSize == 0 && wp.TotalWorker > 3 {
			fmt.Println("Decreasing worker as demand is low")
			wp.Resize(wp.TotalWorker - 1)
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	// create a pool of majdoors with initial majdoors and retry karne ka limit
	pool := newWorkerPool(4, 2)

	// majduro ko kaam karne ko bol do
	pool.Start()

	// actual kaam unko de do
	for i := 1; i < 5; i++ {
		taskID := 1
		task := Task{
			ID: taskID,
			Execute: func() error {
				fmt.Printf("Executing task %d", taskID)
				if taskID%2 == 0 {
					return fmt.Errorf("Task %d failed", taskID)
				}
				return nil
			},
			Retry: 0,
		}
		pool.AddTask(task)
	}
	// stop for 2 seconds for each task kamjor majdoor h
	time.Sleep(2 * time.Second)

	// Resize the worker pool
	fmt.Print("Resizing the worker pool to 7 (Thala for a reason)")
	pool.Resize(7)

	// Nakli kaam ka nakli wait samay
	time.Sleep(5 * time.Second)

	fmt.Println("Stopping worker pool")
	pool.Stop()
}
