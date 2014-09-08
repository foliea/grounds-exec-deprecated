package runner

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"
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

var (
	ErrorTimeoutExceeded = errors.New("Timeout exceeded.")
	ExecutionTimeout     = 10 * time.Second
)

func (r *Runner) Watch() {
	var (
		containerID string
		stop        chan bool
		err         error
	)
	for input := range r.Input {
		var config RunConfig
		if err := json.Unmarshal(input, &config); err != nil {
			r.notifyError(err)
			continue
		}
		if containerID != "" {
			go r.stop(containerID, stop)
		}
		containerID, err = r.Client.Prepare(config.Language, config.Code)
		if err != nil {
			r.notifyError(err)
			continue
		}
		stop = make(chan bool, 3)
		finished := make(chan bool, 1)
		go r.timeout(containerID, finished, stop)
		go r.execute(containerID, finished, stop)
	}
}

func (r *Runner) execute(containerID string, finished, stop chan bool) {
	defer func() {
		if err := r.Client.Clean(containerID); err != nil {
			r.Errs <- err
		}
	}()
	status, err := r.Client.Execute(containerID, func(stdout, stderr io.Reader) {
		go r.broadcast("stdout", stdout, stop)
		r.broadcast("stderr", stderr, stop)
	})
	if err != nil {
		r.notifyError(err)
		return
	}
	finished <- true
	select {
	case <-stop:
	default:
		if status >= 128 {
			status -= 256
		}
		r.write("status", strconv.Itoa(status))
	}
}

func (r *Runner) stop(containerID string, stop chan bool) {
	for i := 0; i < 3; i++ {
		stop <- true
	}
	r.Client.Interrupt(containerID)
}

func (r *Runner) timeout(containerID string, finished, stop chan bool) {
	select {
	case <-finished:
		return
	case <-time.After(ExecutionTimeout):
		r.stop(containerID, stop)
		r.notifyError(ErrorTimeoutExceeded)
	}
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

func (r *Runner) notifyError(err error) {
	r.Errs <- err
	if err == ErrorProgramTooLarge || err == ErrorLanguageNotSpecified || err == ErrorTimeoutExceeded {
		r.write("error", err.Error())
	} else {
		r.write("error", "An error occured.")
	}
}
