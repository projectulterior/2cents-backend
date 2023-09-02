package middleware

import "net/http"

type Cors struct {
	next http.Handler
}

func NewCors(next http.Handler) *Cors {
	return &Cors{next: next}
}

func (c *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	protocol := r.Header.Get("Sec-Websocket-Protocol")
	if protocol == "graphql-ws" {
		c.next.ServeHTTP(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	if r.Method == "OPTIONS" {
		return
	}

	c.next.ServeHTTP(w, r)
}
