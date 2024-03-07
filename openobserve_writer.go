package sharp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// NewOpenObserveWriter will write log to openobserve(https://github.com/openobserve/openobserve)
// url is openobserve server addr. example: http://localhost:5080
func NewOpenObserveWriter(url string, org, group string, username, password string, immediately bool) io.Writer {
	if url == "" || org == "" || group == "" {
		return nil
	}
	p := &openObserveWriter{
		addr:     fmt.Sprintf("%s/api/%s/%s/_multi", url, org, group),
		username: username,
		password: password,
		client:   &http.Client{},
	}
	if !immediately {
		p.ch = make(chan []byte, 4096)
		go p.queue()
	}
	return p
}

type openObserveWriter struct {
	addr string // example: http://localhost:5080/api/myorg/stream1/_multi

	username string
	password string

	client *http.Client

	ch chan []byte
}

func (w *openObserveWriter) queue() {
	for data := range w.ch {
		w.write(data)
	}
}

func (w *openObserveWriter) write(p []byte) (n int, err error) {
	req, err := http.NewRequest(http.MethodPost, w.addr, bytes.NewReader(p))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(w.username, w.password)
	res, err := w.client.Do(req)
	if err != nil {
		return
	}
	res.Body.Close()
	n = len(p)
	return
}

func (w *openObserveWriter) Write(p []byte) (n int, err error) {
	if w.ch != nil {
		n = len(p)
		c := make([]byte, len(p))
		copy(c, p)
		w.ch <- c
		return
	}
	return w.write(p)
}
