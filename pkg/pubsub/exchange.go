package pubsub

import (
	"context"
	"sync"
)

type exchange[M Message] struct {
	ch chan M

	listeners []chan M
	mutex     sync.Mutex
}

func NewExchange[M Message]() Exchange[M] {
	ch := make(chan M)

	e := &exchange[M]{ch: ch}

	go e.run()

	return e
}

func (e *exchange[M]) run() {
	publish := func(msg M) {
		e.mutex.Lock()
		defer e.mutex.Unlock()

		for _, listener := range e.listeners {
			go func(l chan<- M) { l <- msg }(listener)
		}
	}

	for {
		msg := <-e.ch
		publish(msg)
	}
}

func (e *exchange[M]) Publisher() Publisher[M] {
	return NewPublisher[M](e.ch)
}

func (e *exchange[M]) Listener() Listener[M] {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	ch := make(chan M)

	e.listeners = append(e.listeners, ch)

	return &listener[M]{ex: e, ch: ch}
}

func (e *exchange[M]) Subscribe(fn func(context.Context, M) error) {
	ch := make(chan M)

	func() {
		e.mutex.Lock()
		defer e.mutex.Unlock()

		e.listeners = append(e.listeners, ch)
	}()

	for msg := range ch {
		// TODO: handle errors
		_ = fn(context.Background(), msg)
	}
}

func (e *exchange[M]) removeListener(l *listener[M]) {
	l.ex.mutex.Lock()
	defer l.ex.mutex.Unlock()

	for i, listener := range l.ex.listeners {
		if listener == l.ch {
			l.ex.listeners = append(l.ex.listeners[:i], l.ex.listeners[i+1:]...)
			return
		}
	}
}

func (e *exchange[M]) Shutdown(ctx context.Context) error {
	close(e.ch)
	return nil
}
