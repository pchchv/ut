package ut

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	Now       time.Time
	Timestamp int64
	Zone      string
	UTC       time.Time
}

func New() (t Time) {
	t.Now = time.Now()
	t.Timestamp = t.Now.Unix()
	t.UTC = t.Now.UTC()
	zone, _ := t.Now.Zone()
	t.Zone = zone + " " + strings.Split(fmt.Sprint(t.Now), " ")[2]

	return
}
