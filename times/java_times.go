package times

import (
	"time"
)

//TimeToMills CurrentTimeMills
func TimeToMills(current time.Time) int64 {
	return current.UnixNano() / int64(time.Millisecond)
}

//CurrentTimeMills current UNIX
func CurrentTimeMills() int64 {
	return TimeToMills(time.Now())
}
