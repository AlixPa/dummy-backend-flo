package service

import "time"

func Pong() string {
	currentTime := time.Now().Format("15:04:05")
	return "pong at " + currentTime
}
