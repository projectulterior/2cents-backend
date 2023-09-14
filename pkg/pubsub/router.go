package pubsub

import (
	"sync"
)

type router struct {
	ch        chan Message
	listeners []chan<- Message

	mutex sync.Mutex
}

func newRouter() *router {
	ch := make(chan Message)

	r := &router{ch: ch}

	go r.run()

	return r
}

func (r *router) run() {
	publish := func(msg Message) {
		r.mutex.Lock()
		defer r.mutex.Unlock()

		for _, listener := range r.listeners {
			go func(l chan<- Message) { l <- msg }(listener)
		}
	}

	for {
		msg := <-r.ch
		publish(msg)
	}
}

func (r *router) publisher() Publisher {
	return &publisher{ch: r.ch}
}

func (r *router) listener() Listener {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	ch := make(chan Message)

	r.listeners = append(r.listeners, ch)

	return &listener{ch: ch}
}

func (r *router) shutdown() error {
	close(r.ch)
	return nil
}
