package biz

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	loc, _ := time.LoadLocation(timezone)
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), loc)
	*t = LocalTime(now)
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
