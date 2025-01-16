package manager

import (
	"fmt"
	"time"
)

type Worker struct {
	Queue   chan func()
	Workers int
	stop    chan bool
}

func NewWorker(workerCount int) *Worker {
	return &Worker{
		Queue:   make(chan func(), 100),
		Workers: workerCount,
		stop:    make(chan bool),
	}
}

func (w *Worker) Start() {
	fmt.Printf("Worker: Starting %d workers\n", w.Workers)
	for i := 0; i < w.Workers; i++ {
		go func(id int) {
			for {
				select {
				case task := <-w.Queue:
					fmt.Printf("Worker %d: Executing task\n", id)
					task()
				case <-w.stop:
					fmt.Printf("Worker %d: Stopping\n", id)
					return
				}
			}
		}(i + 1)
	}
}

func (w *Worker) AddTask(task func()) {
	w.Queue <- task
	fmt.Println("Worker: Task added to queue")
}

func (w *Worker) Stop() {
	fmt.Println("Worker: Stopping all workers")
	close(w.stop)
	close(w.Queue)
}