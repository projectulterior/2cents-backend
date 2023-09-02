package http

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"

	"go.opencensus.io/zpages"
)

// NewBaseMux creates a ServeMux with various debugging and status endpoints set up:
// /alive, always 200 OK
// /ready, custom handler provided, `httputil.Ready` can be used for this
// /debug/rpcz, RPC Stats
// /debug/tracez, Trace Spans
// /debug/pprof/, pprof
// /debug/pprof/cmdline,  pprof
// /debug/pprof/profile, pprof
// /debug/pprof/symbol, pprof
// /debug/pprof/trace, pprof
func NewBaseMux(ready http.HandlerFunc) *http.ServeMux {
	mux := http.NewServeMux()

	// /alive always responds with 200 OK
	mux.HandleFunc("/alive", func(wr http.ResponseWriter, req *http.Request) {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusOK)
		fmt.Fprint(wr, "OK")
	})

	// /ready is a custom handler, `httputil.Ready` can be used for this
	mux.Handle("/health", ready)

	// zPages exposes various debugging data from OpenCensus
	// endpoints: /debug/rpcz, /debug/tracez
	// more info: https://opencensus.io/zpages/go/
	zpages.Handle(mux, "/debug")

	// pprof allows remote profiling
	// more info: https://golang.org/pkg/net/http/pprof/
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return mux
}

type Server struct {
	Port    int `default:"8080"`
	Handler http.Handler

	wg   sync.WaitGroup
	done chan struct{}
}

func (s *Server) Run() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Handler,
	}

	s.done = make(chan struct{})

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		<-s.done

		if err := server.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Shutdown() {
	close(s.done)
	s.wg.Wait()
}
