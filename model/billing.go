package model

type Billing struct {
	ID          int64  `db:"id"`
	Sender      string `db:"sender"`
	Reciever    string `db:"reciever"`
	Amount      int64  `db:"amount"`
	TimeBill    string `db:"time_bill"`
	TaskID      int64  `db:"task_id"`
	BillingType string `db:"btype"`
}
