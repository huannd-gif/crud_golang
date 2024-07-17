package domain

import (
	"time"
)

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
	Metadata    map[string]interface{}
}

func NewCallAllArgument(
	Id int,
	PhoneNumber string,
	Result string,
	CreateAt *time.Time,
	UpdateAt *time.Time,
	CallAt *time.Time,
	EndAt *time.Time,
	CallPress *time.Time,
	ReceiverAt *time.Time,
	Metadata map[string]interface{},
) *Call {
	return &Call{
		Id:          Id,
		PhoneNumber: PhoneNumber,
		Result:      Result,
		CreateAt:    CreateAt,
		UpdateAt:    UpdateAt,
		CallAt:      CallAt,
		EndAt:       EndAt,
		CallPress:   CallPress,
		ReceiverAt:  ReceiverAt,
		Metadata:    Metadata,
	}
}

func NewCallNoArgument() *Call {
	return &Call{}
}

func (c *Call) GetId() int {
	return c.Id
}

func (c *Call) GetPhoneNumber() string {
	return c.PhoneNumber
}

func (c *Call) GetResult() string {
	return c.Result
}

func (c *Call) GetCreateAt() *time.Time {
	return c.CreateAt
}

func (c *Call) GetUpdateAt() *time.Time {
	return c.UpdateAt
}

func (c *Call) GetCallAt() *time.Time {
	return c.CallAt
}

func (c *Call) GetEndAt() *time.Time {
	return c.EndAt
}

func (c *Call) GetCallPress() *time.Time {
	return c.CallPress
}

func (c *Call) GetReceiverAt() *time.Time {
	return c.ReceiverAt
}

func (c *Call) GetMetadata() map[string]interface{} {
	return c.Metadata
}

// Setter methods
func (c *Call) SetId(id int) {
	c.Id = id
}

func (c *Call) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *Call) SetResult(result string) {
	c.Result = result
}

func (c *Call) SetCreateAt(createAt *time.Time) {
	c.CreateAt = createAt
}

func (c *Call) SetUpdateAt(updateAt *time.Time) {
	c.UpdateAt = updateAt
}

func (c *Call) SetCallAt(callAt *time.Time) {
	c.CallAt = callAt
}

func (c *Call) SetEndAt(endAt *time.Time) {
	c.EndAt = endAt
}

func (c *Call) SetCallPress(callPress *time.Time) {
	c.CallPress = callPress
}

func (c *Call) SetReceiverAt(receiverAt *time.Time) {
	c.ReceiverAt = receiverAt
}

func (c *Call) SetMetadata(metadata map[string]interface{}) {
	c.Metadata = metadata
}
