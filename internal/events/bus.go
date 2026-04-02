package events

import (
	"log"
	"sync"
)

// Event is a named payload that can carry any data.
type Event struct {
	Name    string
	Payload any
}

// Listener is a function that handles an event.
type Listener func(event Event)

// Bus is a simple synchronous pub/sub dispatcher.
// For async dispatch, wrap Listener calls in goroutines.
type Bus struct {
	mu        sync.RWMutex
	listeners map[string][]Listener
}

var defaultBus = &Bus{listeners: make(map[string][]Listener)}

// Listen registers a listener for an event name.
func Listen(eventName string, l Listener) {
	defaultBus.mu.Lock()
	defer defaultBus.mu.Unlock()
	defaultBus.listeners[eventName] = append(defaultBus.listeners[eventName], l)
}

// Dispatch fires an event to all registered listeners.
func Dispatch(eventName string, payload any) {
	defaultBus.mu.RLock()
	listeners := defaultBus.listeners[eventName]
	defaultBus.mu.RUnlock()

	e := Event{Name: eventName, Payload: payload}
	for _, l := range listeners {
		l(e)
	}
}

// DispatchAsync fires the event in a goroutine (fire-and-forget).
func DispatchAsync(eventName string, payload any) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("events: panic in listener for '%s': %v", eventName, r)
			}
		}()
		Dispatch(eventName, payload)
	}()
}

// Built-in event names — add yours here to avoid magic strings.
const (
	UserCreated = "user.created"
	UserDeleted = "user.deleted"
)
