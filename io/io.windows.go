//go:build windows
// +build windows

package io

func ReadFileToBytes(src string) (rs []byte, err error) {
	return rs, err
}

func SetMaxRlimit() error {
	return nil
}
