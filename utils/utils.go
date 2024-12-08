package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func SplitPkgVersion(s string) (pkg, ver string) {
	splits := strings.Split(s, "@")
	if len(splits) < 2 {
		return splits[0], ""
	}
	pkg, ver = splits[0], splits[1]
	return
}

func GetImageFormat(s string) (hasFormat bool, format string) {
	splits := strings.Split(s, ".")
	if len(splits) < 2 {
		return false, "svg"
	}
	return true, splits[len(splits)-1]
}

func GetInputScanner(inputFileName string) (*bufio.Scanner, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if inputFileName == "" && fi.Mode()&os.ModeNamedPipe != os.ModeNamedPipe {
		return nil, errors.New("not enough inputs")
	}
	var s *bufio.Scanner
	if fi.Mode()&os.ModeNamedPipe != os.ModeNamedPipe {
		inputFile, err := os.Open(inputFileName)
		if err != nil {
			return nil, err
		}
		s = bufio.NewScanner(inputFile)
	} else {
		s = bufio.NewScanner(os.Stdin)
	}
	return s, nil
}
