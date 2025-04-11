package ping

import "time"

func GetPingResponse() string {
	currentTime := time.Now().Format("15:04:05")
	return "pong at " + currentTime
}
