package utils

import "fmt"

func Must[T any](err error) {
	if err != nil {
		panic(fmt.Sprintf("err must be nil, but got: %+v", err))
	}
}

func Must2[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("err must be nil, but got: %+v", err))
	}

	return val
}
