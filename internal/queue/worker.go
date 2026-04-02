package queue

import (
	"context"
	"encoding/json"
	"log"

	goredis "github.com/redis/go-redis/v9"
)

const defaultQueue = "go-skeleton:jobs"

// Job represents a unit of async work pushed to Redis.
type Job struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// Handler is the function signature for job processors.
type Handler func(ctx context.Context, payload json.RawMessage) error

// Worker listens to a Redis list and dispatches jobs to registered handlers.
type Worker struct {
	rdb      *goredis.Client
	handlers map[string]Handler
	queue    string
}

func NewWorker(rdb *goredis.Client) *Worker {
	return &Worker{
		rdb:      rdb,
		handlers: make(map[string]Handler),
		queue:    defaultQueue,
	}
}

// Register binds a job type string to a handler function.
func (w *Worker) Register(jobType string, h Handler) {
	w.handlers[jobType] = h
}

// Run blocks and processes jobs until ctx is cancelled.
func (w *Worker) Run(ctx context.Context) {
	log.Printf("🔄 Queue worker started — listening on '%s'", w.queue)
	for {
		select {
		case <-ctx.Done():
			log.Println("Queue worker stopped")
			return
		default:
			result, err := w.rdb.BLPop(ctx, 0, w.queue).Result()
			if err != nil {
				continue
			}
			raw := result[1]
			var job Job
			if err := json.Unmarshal([]byte(raw), &job); err != nil {
				log.Printf("queue: malformed job: %v", err)
				continue
			}
			h, ok := w.handlers[job.Type]
			if !ok {
				log.Printf("queue: no handler for job type '%s'", job.Type)
				continue
			}
			if err := h(ctx, job.Payload); err != nil {
				log.Printf("queue: job '%s' failed: %v", job.Type, err)
			}
		}
	}
}

// Dispatch pushes a job onto the Redis queue.
func Dispatch(ctx context.Context, rdb *goredis.Client, jobType string, payload any) error {
	p, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	job := Job{Type: jobType, Payload: p}
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return rdb.RPush(ctx, defaultQueue, data).Err()
}
