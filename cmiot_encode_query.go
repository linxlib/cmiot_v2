package cmiot_v2

import (
	"net/http"
	"net/url"
	"reflect"
)

var _ Adder = (*QueryEncode)(nil)

// QueryEncode URL query encoder structure
type QueryEncode struct {
	values url.Values
	r      *http.Request
}

// newQueryEncode create a new URL query  encoder
func newQueryEncode(req *http.Request) *QueryEncode {
	return &QueryEncode{values: make(url.Values)}
}

// add Encoder core function, used to set each key / value into the http URL query
func (q *QueryEncode) add(key string, v reflect.Value, sf reflect.StructField) error {
	q.values.Add(key, valToStr(v, sf))
	return nil
}

// end URL query structured data into strings
func (q *QueryEncode) end() string {
	return q.values.Encode()
}

// name URL query Encoder name
func (q *QueryEncode) name() string {
	return "query"
}
