package runner

import (
	"bufio"
	"encoding/json"
	"io"
	"strconv"
)

type RunConfig struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Response struct {
	Stream string `json:"stream"`
	Chunk  string `json:"chunk"`
}

type Runner struct {
	Client *Client

	Input  chan []byte
	Output chan []byte
	Errs   chan error
}

func (r *Runner) Read() {
	var (
		containerID string
		stop        chan bool
		err         error
	)
	for input := range r.Input {
		var config RunConfig
		json.Unmarshal(input, &config)

		if containerID != "" {
			go r.stop(containerID, stop)
		}

		containerID, err = r.Client.Prepare(config.Language, config.Code)
		if err != nil {
			r.Errs <- err
			continue
		}
		stop = make(chan bool, 1)
		go r.execute(containerID, stop)
	}
	close(r.Output)
	close(r.Errs)
}

func (r *Runner) execute(containerID string, stop chan bool) {
	defer r.Client.Clean(containerID)
	status, err := r.Client.Execute(containerID, func(stdout, stderr io.Reader) {
		go r.broadcast("stdout", stdout, stop)
		r.broadcast("stderr", stderr, stop)
	})
	if err != nil {
		r.Errs <- err
		return
	}
	select {
	case <-stop:
	default:
		r.write("status", strconv.Itoa(status))
	}
}

func (r *Runner) stop(containerID string, stop chan bool) {
	for i := 0; i < 3; i++ {
		stop <- true
	}
	r.Client.Interrupt(containerID)
}

func (r *Runner) write(stream, chunk string) {
	response, err := json.Marshal(Response{Stream: stream, Chunk: chunk})
	if err != nil {
		r.Errs <- err
		return
	}
	r.Output <- response
}

func (r *Runner) broadcast(stream string, output io.Reader, stop chan bool) {
	var (
		reader = bufio.NewReader(output)
		buffer = make([]byte, 1024)
	)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return
		}
		select {
		case <-stop:
			return
		default:
			if n > 0 {
				r.write(stream, string(buffer[0:n]))
			}
		}
	}
}
