package mocks

import "io"

type Streams struct {
	StdinF  func() io.Reader
	StdoutF func() io.Writer
	StderrF func() io.Writer
}

func (s *Streams) Stdin() io.Reader {
	return s.StdinF()
}

func (s *Streams) Stdout() io.Writer {
	return s.StdoutF()
}

func (s *Streams) Stderr() io.Writer {
	return s.StderrF()
}
