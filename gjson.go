//go:build 386 || arm || mipsle || mips || solaris || illumos || plan9

package json

import (
	"github.com/goccy/go-json"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
