package monitor

import (
	"fmt"
	"strings"
	"time"
)

// MyTime wraps the time.Time object so
// that it can be Unmarshalled/marshalled
// from a simpler layout 2006/01/02 ...
type MyTime struct {
	time.Time
}

// UnmarshalJSON let MyTime be Unmarshalled in a flexible way.
func (mt *MyTime) UnmarshalJSON(b []byte) (err error) {
	// Trim string aggressively
	s := strings.Trim(string(b), "\"\n\r ")
	if s == "null" {
		mt.Time = time.Time{}
		return nil
	}
	// Try various format until one matches ...
	mt.Time, err = time.Parse("2006-01-02", s)
	if err == nil {
		return
	}
	mt.Time, err = time.Parse("2006-1-2", s)
	if err == nil {
		return
	}
	mt.Time, err = time.Parse("2006-01-02 15:04:05 -0700 MST", s)
	if err == nil {
		return
	}
	mt.Time, err = time.Parse(time.RFC3339, s)
	if err == nil {
		return
	}
	mt.Time, err = time.Parse(time.RFC822, s)
	if err == nil {
		return
	}
	mt.Time, err = time.Parse(time.RFC822Z, s)
	if err == nil {
		return
	}
	if err != nil {
		fmt.Println("Could not parse : ", mt)
	}
	return
}

// MarshalJSON in simplified format.
func (mt *MyTime) MarshalJSON() ([]byte, error) {
	if mt.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", mt.Time.Format("2006-01-02"))), nil
}

// zero value
var nilTime = (time.Time{}).UnixNano()

// IsSet check for zero value.
func (mt *MyTime) IsSet() bool {
	return mt.UnixNano() != nilTime
}

func (mt *MyTime) String() string {
	return mt.Time.Format("2006-01-02")
}
