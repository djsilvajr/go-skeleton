package scheduler

import (
	"context"
	"log"
	"time"
)

// Task is a scheduled function with a name and interval.
type Task struct {
	Name     string
	Interval time.Duration
	Run      func(ctx context.Context)
}

// Scheduler runs registered tasks on their intervals.
type Scheduler struct {
	tasks []Task
}

func New() *Scheduler {
	return &Scheduler{}
}

// Add registers a task.
func (s *Scheduler) Add(name string, interval time.Duration, fn func(ctx context.Context)) {
	s.tasks = append(s.tasks, Task{Name: name, Interval: interval, Run: fn})
}

// Start launches all tasks in goroutines and blocks until ctx is cancelled.
func (s *Scheduler) Start(ctx context.Context) {
	log.Println("⏰ Scheduler started")
	for _, task := range s.tasks {
		t := task
		go func() {
			ticker := time.NewTicker(t.Interval)
			defer ticker.Stop()
			log.Printf("scheduler: task '%s' every %s", t.Name, t.Interval)
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					log.Printf("scheduler: running '%s'", t.Name)
					t.Run(ctx)
				}
			}
		}()
	}
	<-ctx.Done()
	log.Println("Scheduler stopped")
}
