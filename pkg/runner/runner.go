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
		if err := json.Unmarshal(input, &config); err != nil {
			r.handleError(err)
			continue
		}
		if containerID != "" {
			go r.stop(containerID, stop)
		}
		containerID, err = r.Client.Prepare(config.Language, config.Code)
		if err != nil {
			r.handleError(err)
			continue
		}
		stop = make(chan bool, 3)
		go r.execute(containerID, stop)
	}
	r.stop(containerID, stop)
	close(r.Output)
	close(r.Errs)
}

func (r *Runner) execute(containerID string, stop chan bool) {
	defer func() {
		if err := r.Client.Clean(containerID); err != nil {
			r.handleError(err)
		}
	}()
	status, err := r.Client.Execute(containerID, func(stdout, stderr io.Reader) {
		go r.broadcast("stdout", stdout, stop)
		r.broadcast("stderr", stderr, stop)
	})
	if err != nil {
		r.handleError(err)
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
		r.handleError(err)
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

func (r *Runner) handleError(err error) {
	r.Errs <- err
	if err == ErrorProgramTooLarge || err == ErrorLanguageNotSpecified {
		r.write("error", err.Error())
	}
}
