package pubsub

import (
	"sync"
)

type router struct {
	ch        chan Message
	listeners []chan Message

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

	return &listener{router: r, ch: ch}
}

func (r *router) removeListener(l *listener) {
	l.router.mutex.Lock()
	defer l.router.mutex.Unlock()

	for i, listener := range l.router.listeners {
		if listener == l.ch {
			l.router.listeners = append(l.router.listeners[:i], l.router.listeners[i+1:]...)
			return
		}
	}
}

func (r *router) shutdown() error {
	close(r.ch)
	return nil
}
