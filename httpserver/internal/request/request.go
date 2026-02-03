package request

import (
	"bytes"
	"fmt"
	"io"
)

type Request struct {
	RequestLine  RequestLine
	RequestState RequestState
}

type RequestState = int

const (
	RequestStateInit = 0
	RequestStateDone = 1
)

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

func (rq *Request) parse(b []byte) (int, error) {
	read := 0
outer:
	for {
		switch rq.RequestState {
		case RequestStateInit:
			rl, n, err := parseRequestLine(b)
			if err != nil {
				return 0, err
			}
			if n == 0 {
				break outer
			}

			read += n
			rq.RequestLine = rl

			rq.RequestState = RequestStateDone

		case RequestStateDone:
			return read, nil
		}
	}

	return 0, nil
}

func (rq *Request) done() bool {
	return rq.RequestState == RequestStateDone
}

func RequestFromReader(r io.Reader) (*Request, error) {
	request := &Request{RequestState: RequestStateInit}

	buf := make([]byte, 1024)
	buf_len := 0

	for !request.done() {
		n, err := r.Read(buf[buf_len:])
		if err != nil {
			return nil, err
		}
		if n == 0 {
			continue
		}

		buf_len += n

		parse_n, err := request.parse(buf[:buf_len+n])
		if err != nil {
			return nil, err
		}

		if parse_n > 0 {
			copy(buf, buf[parse_n:buf_len])
			buf_len -= parse_n
		}
	}

	return request, nil
}

var SEPERATOR = []byte("\r\n")
var ERROR_BAD_RL error = fmt.Errorf("Error: Bad Request Line")

func parseRequestLine(b []byte) (RequestLine, int, error) {
	rq := RequestLine{}
	idx := bytes.Index(b, SEPERATOR)
	if idx == -1 {
		return rq, 0, nil
	}

	parts := bytes.Split(b[:idx], []byte(" "))
	if len(parts) != 3 {
		return rq, 0, ERROR_BAD_RL
	}

	for _, c := range parts[0] {
		if c < 65 || c > 90 {
			return rq, 0, ERROR_BAD_RL
		}
	}

	var version string
	if string(parts[2]) != "HTTP/1.1" {
		return rq, 0, ERROR_BAD_RL
	} else {
		version = "1.1"
	}

	return RequestLine{Method: string(parts[0]), RequestTarget: string(parts[1]), HttpVersion: version}, idx + len(SEPERATOR), nil
}
