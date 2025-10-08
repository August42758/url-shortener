package rest

import (
	"fmt"
)

func catchPanic(funcName string) {
	p := recover()
	if p != nil {
		fmt.Printf("В %v произошла паника: %v\n", funcName, p)
	}
}
