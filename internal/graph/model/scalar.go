package model

import (
	"errors"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// type Timestamp int64

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		log.Print(t.Unix())
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
