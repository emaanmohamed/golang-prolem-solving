package main

import (
	"fmt"
	"sync"
	"time"
)

// // Implement a concurrent worker pool in Go.
// Task is a function type that represents a task to be executed by a worker.

type Task func()

// Worker represents a worker that processes tasks from the task channel.
type Worker struct {
	id       int
	taskChan <-chan Task
	wg       *sync.WaitGroup
}

// NewWorker creates a new Worker.
func NewWorker(id int, taskChan <-chan Task, wg *sync.WaitGroup) Worker {
	return Worker{
		id:       id,
		taskChan: taskChan,
		wg:       wg,
	}
}

// Start begins the worker's task processing loop.
func (w Worker) Start() {
	go func() {
		for task := range w.taskChan {
			fmt.Printf("Worker %d started task\n", w.id)
			task()
			fmt.Printf("Worker %d finished task\n", w.id)
			w.wg.Done()
		}
	}()
}

// WorkerPool represents a pool of workers.
type WorkerPool struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with a specified number of workers.
func NewWorkerPool(numWorkers int) *WorkerPool {
	taskChan := make(chan Task)
	pool := &WorkerPool{
		taskChan: taskChan,
	}
	pool.startWorkers(numWorkers)
	return pool
}

// startWorkers initializes the specified number of workers.
func (p *WorkerPool) startWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i+1, p.taskChan, &p.wg)
		worker.Start()
	}
}

// AddTask adds a new task to the worker pool.
func (p *WorkerPool) AddTask(task Task) {
	p.wg.Add(1)
	p.taskChan <- task
}

// Wait waits for all tasks to be completed.
func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.taskChan)
}

func main() {
	// Create a worker pool with 3 workers.
	pool := NewWorkerPool(3)

	// Define some tasks.
	tasks := []Task{
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 1 completed")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Task 2 completed")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("Task 3 completed")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Task 4 completed")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 5 completed")
		},
	}

	// Add tasks to the pool.
	for _, task := range tasks {
		pool.AddTask(task)
	}

	// Wait for all tasks to complete.
	pool.Wait()
	fmt.Println("All tasks completed")
}
