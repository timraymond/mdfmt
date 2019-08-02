package main

import (
	"io"
	"os"

	"github.com/timraymond/mdfmt/fmt"
)

type OSStreams struct {
}

func (o *OSStreams) Stdout() io.Writer {
	return os.Stdout
}

func (o *OSStreams) Stdin() io.Reader {
	return os.Stdin
}

func (o *OSStreams) Stderr() io.Writer {
	return os.Stderr
}

func main() {
	c := fmt.Command{
		Streams: &OSStreams{},
	}
	c.Run()
}
