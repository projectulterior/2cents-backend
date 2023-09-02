package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type requestLogger struct {
	log  *zap.Logger
	next http.Handler
}

func NewRequestLogger(logger *zap.Logger, next http.Handler) http.Handler {
	return &requestLogger{
		log:  logger,
		next: next,
	}
}

func (l *requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	r.Body = NewTeeReadCloser(r.Body, &buf)

	response := &loggerResponseWriter{ResponseWriter: w}

	start := time.Now()
	defer func() {
		duration := time.Since(start)

		req, err := ioutil.ReadAll(&buf)
		if err != nil {
			req = []byte("error in reading request")
		}

		l.log.Info(fmt.Sprintf("[%d] %s %s %s", response.code, r.Method, r.RequestURI, duration),
			zap.String("request", string(req)),
			zap.Duration("dur", duration),
		)
	}()

	l.next.ServeHTTP(response, r)
}

type teeReadCloser struct {
	r io.Reader
	c io.Closer
}

func NewTeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	return &teeReadCloser{
		r: io.TeeReader(r, w),
		c: r,
	}
}

func (t *teeReadCloser) Read(p []byte) (int, error) {
	return t.r.Read(p)
}

func (t *teeReadCloser) Close() error {
	return t.c.Close()
}
