package query

import "time"

type Call struct {
	Id          int
	PhoneNumber string
	Result      string
	CreateAt    *time.Time
	UpdateAt    *time.Time
	CallAt      *time.Time
	EndAt       *time.Time
	CallPress   *time.Time
	ReceiverAt  *time.Time
}

type Paging struct {
	PageNum    int
	PageSize   int
	PageTotal  int
	TotalCount int
}

type ListCallsPaginated struct {
	Calls    []*Call
	Metadata Paging
}
