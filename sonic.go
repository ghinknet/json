//go:build !386 && !arm && !mipsle && !mips && !solaris && !illumos && !plan9

package json

import "github.com/bytedance/sonic"

func Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}
