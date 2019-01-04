package pkg

import (
	"fmt"
	"strconv"
	"time"
)

type UTCTimestamp struct {
	time.Time
}

func (t *UTCTimestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(t.Time).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *UTCTimestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = UTCTimestamp{time.Unix(int64(ts), 0)}

	return nil
}

func (t *UTCTimestamp) String() string {
	return fmt.Sprint(time.Time(t.Time).Unix())
}
