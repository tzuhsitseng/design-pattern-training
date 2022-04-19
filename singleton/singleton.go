package main

import (
	"fmt"
	"sync"
)

var (
	once         sync.Once
	countryCodes map[string]string
)

func init() {
	once.Do(func() {
		countryCodes = map[string]string{
			"AF": "004",
			"TW": "158",
			"US": "840",
		}
	})
}

func main() {
	for k, v := range countryCodes {
		fmt.Printf("%s: %s\n", k, v)
	}
}
