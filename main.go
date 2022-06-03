package main

import (
	"fmt"
	"mido/messages"
)

func main() {
	fmt.Println("SPECS")
	for _, v := range messages.SPECS {
		fmt.Println(v)
	}
	fmt.Println("SPEC_LOOKUP")
	for k, v := range messages.SPEC_LOOKUP {
		fmt.Println(k, " ", v)
	}
	fmt.Println("SPEC_BY_STATUS")
	for k, v := range messages.SPEC_BY_STATUS {
		fmt.Println(k, " ", v)
	}
	fmt.Println("SPEC_BY_TYPE")
	for k, v := range messages.SPEC_BY_TYPE {
		fmt.Println(k, " ", v)
	}
}
