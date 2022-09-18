//go:build !windows
// +build !windows

package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

func ReadFileToBytes(src string) (rs []byte, err error) {
	sourceFileStat, err := os.Stat(filepath.Clean(src))
	if err != nil {
		return rs, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return rs, fmt.Errorf("%s is not a regular file", filepath.Clean(src))
	}
	source, err := os.Open(filepath.Clean(src))
	if err != nil {
		return rs, err
	}

	dest := bytes.NewBuffer(nil)
	_, err = io.Copy(dest, source)
	if err != nil {
		return rs, err
	}

	err = source.Close()

	return dest.Bytes(), err
}

func SetMaxRlimit() error {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}
	rLimit.Cur = rLimit.Max
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
