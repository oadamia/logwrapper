package logwrapper

import "time"

func utcTimeFunc() time.Time {
	return time.Now().UTC()
}
