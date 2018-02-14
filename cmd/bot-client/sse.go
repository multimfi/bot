package main

import (
	"bytes"
	"errors"
	"strconv"
)

func peek(data []byte, n int, r byte) bool {
	if len(data) <= n {
		return false
	}
	return data[n] == r
}

func seq(data []byte) int {
	var n int
	for _, v := range data {
		if v != '\n' {
			break
		}
		n++
	}
	return n
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	var n int
	for {
		i := bytes.IndexByte(data[n:], '\n')
		if i < 0 {
			break
		}
		if peek(data, n+i+1, '\n') {
			x := seq(data[n+i+1:])
			return n + i + 1 + x, data[0 : n+i], nil
		}
		n += i + 1
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

var errInvalidEvent = errors.New("sse: invalid event")

type SSE struct {
	id    int
	event int
	data  []byte
}

func parseSSE(data []byte) (SSE, error) {
	var ret SSE

	r := bytes.SplitN(data, []byte{'\n'}, 3)
	for _, v := range r {
		i := bytes.IndexByte(v, ':')
		if i < 0 {
			continue
		}

		switch string(v[:i]) {
		case "id":
			x, err := strconv.Atoi(string(v[i+1:]))
			if err != nil {
				return ret, errInvalidEvent
			}
			ret.id = x

		case "event":
			x, err := strconv.Atoi(string(v[i+1:]))
			if err != nil {
				return ret, errInvalidEvent
			}
			ret.event = x

		case "data":
			ret.data = v[i+1:]

		default:
			return ret, errInvalidEvent
		}
	}

	return ret, nil
}
