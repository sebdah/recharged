package types

import (
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

// Override the Time.MarshalJSON function
func (j *JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", j.Time.UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}

// Override the Time.String function
func (j *JSONTime) String() string {
	return j.Time.UTC().Format(time.RFC3339)
}

// Override the Time.UnmarshalJSON function
func (j *JSONTime) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	j.Time, err = time.Parse(fmt.Sprintf("\"%s\"", time.RFC3339), string(data))
	return
}
