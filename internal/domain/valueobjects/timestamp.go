package valueobjects 

import(
	"time"
	"database/sql/driver"
)
type Timestamp time.Time

func(t *Timestamp) Scan(val interface{} ) error {
	*t = Timestamp(val.(time.Time))
	return nil
}

func(t *Timestamp) Value() driver.Value, error {
	return time.Time(t),nil
}