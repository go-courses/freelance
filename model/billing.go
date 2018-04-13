package model

import "time"

type Billing struct {
	ID          int64     `db:"id"`
	Sender      string    `db:"sender"`
	Reciever    string    `db:"reciever"`
	Amount      int       `db:"amount"`
	TimeBill    time.Time `db:"time_bill"`
	TaskID      int64     `db:"task_id"`
	BillingType string    `db:"btype"`
}
