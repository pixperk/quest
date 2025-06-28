package ui

import "time"

type ResponseMessage struct {
	StatusCode   int
	Headers      map[string]string
	Body         string
	ResponseTime time.Duration
	Error        error
}
