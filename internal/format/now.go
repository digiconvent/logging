package format

import "time"

func now() string {
	return time.Now().Format("2006.01.02T15:04:05")
}
