package ut

import (
	"time"
)

type Time struct {
	Now       time.Time
	Timestamp int64
	Zone      string
	UTC       time.Time
}
