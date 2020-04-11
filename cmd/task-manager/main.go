package main

import (
	"fmt"

	timers "github.com/chukak/task-manager/pkg/timers"
)

func main() {
	// todo
	fmt.Println("Hello World!")

	t := timers.NewElapsedTimer()
	_ = t
}
