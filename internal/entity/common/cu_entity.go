package common

import "time"

type CUEntity struct {
	CreateDate *time.Time `json:"create_date"`
	CreateId   int        `json:"create_id"`
	UpdateDate *time.Time `json:"update_date"`
	UpdateId   int        `json:"update_id"`
}
