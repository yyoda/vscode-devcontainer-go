package model

import (
	"time"
)

// Timestamp .
type Timestamp struct {
	Created time.Time `db:"Created"`
	Updated time.Time `db:"Updated"`
}
