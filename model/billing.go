package model

import "time"

// Billing ...
type Billing struct {
	ID          int64     `db:"id"`
	Sender      int64     `db:"sender"`
	Reciever    int64     `db:"reciever"`
	Amount      Money     `db:"amount"`
	TimeBill    time.Time `db:"time_bill"`
	TaskID      int32     `db:"task_id"`
	BillingType string    `db:"btype"`
}
