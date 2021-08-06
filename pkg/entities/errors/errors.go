package errors

import "errors"

//goland:noinspection ALL
var MalformedGraph = errors.New("malformed graph")
var UnknownHttpResponse = errors.New("unknown http response")

type HttpError struct {
	Message string
}
