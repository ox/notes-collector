package notefile

import (
	"time"
)

type Note struct {
	Timestamp time.Time
	Link      string
	Text      string
}
