package resolver

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const layout = "2006-01-02T15:04:05.000Z07:00"

func MarshalTime(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(layout)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(layout, tmpStr)
	}
	return time.Time{}, fmt.Errorf("time should be formated as: %s", layout)
}
