package broker

import (
	"sync"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
)

var exchanges = make(map[string]interface{})
var mutex sync.Mutex

func Exchange[M pubsub.Message](m M) pubsub.Exchange[M] {
	mutex.Lock()
	defer mutex.Unlock()

	ex, ok := exchanges[m.Route()].(pubsub.Exchange[M])
	if !ok {
		ex = pubsub.NewExchange[M]()
		exchanges[m.Route()] = ex
	}

	return ex
}
