package models

import (
	"time"
)

type DepositItem struct {

	// item to be recyled eg. small-bottle
	ItemSize string `json:"itemSize,omitempty"`

	// date and time item was recycled
	DateTime time.Time `json:"dateTime,omitempty"`
}
