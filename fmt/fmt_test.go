package fmt_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/matryer/is"

	"github.com/timraymond/mdfmt/fmt"
	"github.com/timraymond/mdfmt/mocks"
)

func Test_Format(t *testing.T) {
	is := is.New(t)

	sub, err := os.Open("./testdata/raw.md")
	is.NoErr(err) // opening sub

	got := bytes.NewBufferString("")
	errs := bytes.NewBufferString("")

	expFile, err := os.Open("./testdata/exp.md")
	is.NoErr(err) // opening exp

	exp, err := ioutil.ReadAll(expFile)
	is.NoErr(err) // reading exp

	c := &fmt.Command{
		Streams: &mocks.Streams{
			StdinF: func() io.Reader {
				return sub
			},
			StdoutF: func() io.Writer {
				return got
			},
			StderrF: func() io.Writer {
				return errs
			},
		},
	}

	if errs.Len() > 0 {
		t.Log(string(errs.Bytes()))
		t.FailNow()
	}

	err = c.Run()
	is.NoErr(err) // running the command

	if !bytes.Equal(exp, got.Bytes()) {
		t.Log("exp does not match got:")
		t.Log("exp:")
		t.Log(string(exp))
		t.Log("got:")
		t.Log(got.String())
		t.Fail()
	}
}
