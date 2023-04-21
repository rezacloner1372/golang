package main

import (
	"fmt"
	"runtime"
	"time"
)

func compare(x, y int) (z string) {
	if x > y {
		z = "x is greater than y"
	} else {
		z = "x is smaller than y"
	}
	return
}

func CheckOS() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return os
	case "windows":
		return os
	default:
		return os

	}
}

func TimeChecker() string {
	time := time.Now().Hour()
	switch {
	case time < 12:
		return "Good morning"
	case time < 17:
		return "Good afternoon"
	default:
		return "happy"
	}
}

func main() {
	t := compare(1, 4)

	fmt.Println(t)
	fmt.Println(CheckOS())
	fmt.Println(TimeChecker())
}
