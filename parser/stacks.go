package parser

import (
	"bytes"
	"path/filepath"

	"github.com/DataDog/gostackparse"
)

type Stacks struct {
	Func    string `json:"func"`
	Package string `json:"package"`
	Line    int    `json:"line"`
}

func ParseStack(stack []byte) (res []Stacks) {
	if goroutines, err := gostackparse.Parse(bytes.NewReader(stack)); err == nil {
		for _, val := range goroutines {
			for _, stt := range val.Stack {
				if stt != nil {
					res = append(res, Stacks{
						Func:    filepath.Base(stt.Func),
						Package: filepath.Dir(stt.Func),
						Line:    stt.Line,
					})

				}
			}
		}
	}

	if len(res) > 10 {
		res = res[:10]
	}

	return res
}
