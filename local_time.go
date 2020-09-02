package resengo

import (
	"encoding/json"
	"time"
)

type LocalTime struct {
	time.Time
}

func (d *LocalTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05"))
}

func (d LocalTime) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *LocalTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d LocalTime) String() string {
	return d.Time.Format("2006-01-02T15:04:05")
}
