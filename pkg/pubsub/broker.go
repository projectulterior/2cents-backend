package pubsub

import (
	"context"
	"sync"
)

type broker struct {
	mutex   sync.Mutex
	routers map[Route]*router
}

func NewBroker() Broker {
	return &broker{
		routers: make(map[Route]*router),
	}
}

func (b *broker) getRouter(route Route) *router {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	router, ok := b.routers[route]
	if !ok {
		router = newRouter()

		b.routers[route] = router
	}

	return router
}

func (b *broker) Publisher(route Route) Publisher {
	return b.getRouter(route).publisher()
}

func (b *broker) Listener(route Route) Listener {
	return b.getRouter(route).listener()
}

func (b *broker) Shutdown(ctx context.Context) error {
	for _, router := range b.routers {
		if err := router.shutdown(); err != nil {
			return err
		}
	}

	return nil
}
