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
	t.Zone = strings.Split(fmt.Sprint(t.Now), " ")[2]

	return
}

func (t *Time) byLocation(zone string) (inZoneTime time.Time, err error) {
	location, err := time.LoadLocation(zone)
	if err != nil {
		return
	}

	return t.Now.In(location), nil
}
