package sharp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// NewZincsearchWriter will write log to zincsearch(https://github.com/zincsearch/zincsearch)
// url is zincsearch server addr. example: http://localhost:4080
func NewZincsearchWriter(url string, namespace string, username, password string, immediately bool) io.Writer {
	if url == "" || namespace == "" {
		return nil
	}
	p := &zincsearchWriter{
		addr:     fmt.Sprintf("%s/api/%s/_doc", url, namespace),
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

type zincsearchWriter struct {
	addr string // example: http://localhost:4080/api/myapp/_doc

	username string
	password string

	client *http.Client

	ch chan []byte
}

func (w *zincsearchWriter) queue() {
	for data := range w.ch {
		w.write(data)
	}
}

func (w *zincsearchWriter) write(p []byte) (n int, err error) {
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

func (w *zincsearchWriter) Write(p []byte) (n int, err error) {
	if w.ch != nil {
		n = len(p)
		c := make([]byte, len(p))
		copy(c, p)
		w.ch <- c
		return
	}
	return w.write(p)
}
