package fmt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Streams interface {
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
}

// Command represents the environment that a command will run in
type Command struct {
	Streams Streams
}

func (c *Command) Run() error {
	s := bufio.NewScanner(c.Streams.Stdin())
	s.Split(scanSentences)

	w := c.Streams.Stdout()
	e := c.Streams.Stderr()

	for s.Scan() {
		_, err := w.Write(s.Bytes())
		if err != nil {
			e.Write([]byte(fmt.Sprintf("Error writing: err: %v", err)))
		}

		_, err = w.Write([]byte(".\n"))
		if err != nil {
			e.Write([]byte(fmt.Sprintf("Error writing newline: err: %v", err)))
		}
	}

	return nil
}

func scanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '.'); i >= 0 {
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
