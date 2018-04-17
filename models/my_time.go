package models

import (
	"time"
	"strings"
	"fmt"
)

type JsonTime time.Time

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*jt = JsonTime(t)
	return nil
}

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(jt).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Maybe a Format function for printing your date
//func (j JsonTime) Format(s string) string {
//	t := time.Time(j)
//	return t.Format(s)
//}