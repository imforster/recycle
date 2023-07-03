package models

import (
	"time"
)

type DepositItem struct {
	// cubeId
	CubeId string `json:"cubeId,omitempty"`

	// customerId
	CustomerId string `json:"customerId,omitempty"`

	// item to be recyled eg. small-bottle
	ItemSize string `json:"itemSize,omitempty"`

	// date and time item was recycled
	DateTime time.Time `json:"dateTime,omitempty"`
}
