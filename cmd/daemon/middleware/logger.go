package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	next http.Handler
}

func NewLogger(next http.Handler) *Logger {
	return &Logger{next: next}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	protocol := r.Header.Get("Sec-Websocket-Protocol")
	if protocol == "graphql-ws" {
		l.next.ServeHTTP(w, r)
		return
	}

	response := &loggerResponseWriter{ResponseWriter: w, Hijacker: w.(http.Hijacker)}

	start := time.Now()
	defer func() {
		if len(response.err) > 1000 {
			response.err = response.err[:1000]
		}

		log.Printf("[%d] %s %s %s %s", response.code, r.Method, r.RequestURI, time.Since(start), response.err)
	}()

	// req, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(fmt.Sprintf("error in reading body: %s", err)))
	// 	return
	// }
	// log.Printf("[request] %s", string(req))

	l.next.ServeHTTP(response, r)
}

type loggerResponseWriter struct {
	http.ResponseWriter
	http.Hijacker
	code int
	err  []byte
}

func (l *loggerResponseWriter) WriteHeader(code int) {
	l.code = code
	l.ResponseWriter.WriteHeader(code)
}

func (l *loggerResponseWriter) Write(p []byte) (int, error) {
	if l.code < 200 || l.code > 299 {
		l.err = append(l.err, p...)
	}

	return l.ResponseWriter.Write(p)
}
