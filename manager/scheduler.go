package manager

import (
	"fmt"
	"time"
)

type ScheduledTask struct {
	Name     string
	Interval time.Duration
	Task     func()
}

type Scheduler struct {
	Tasks []ScheduledTask
	stop  chan bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Tasks: []ScheduledTask{},
		stop:  make(chan bool),
	}
}

func (s *Scheduler) AddTask(name string, interval time.Duration, task func()) {
	s.Tasks = append(s.Tasks, ScheduledTask{
		Name:     name,
		Interval: interval,
		Task:     task,
	})
	fmt.Printf("Scheduler: Task '%s' added with interval %s\n", name, interval)
}

func (s *Scheduler) Start() {
	fmt.Println("Scheduler: Starting all tasks")
	for _, task := range s.Tasks {
		go func(t ScheduledTask) {
			ticker := time.NewTicker(t.Interval)
			for {
				select {
				case <-ticker.C:
					fmt.Printf("Scheduler: Running task '%s'\n", t.Name)
					t.Task()
				case <-s.stop:
					ticker.Stop()
					return
				}
			}
		}(task)
	}
}

func (s *Scheduler) Stop() {
	fmt.Println("Scheduler: Stopping all tasks")
	close(s.stop)
}