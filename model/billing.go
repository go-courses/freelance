package model

type Billing struct {
	Id           int64
	Sender       string
	Reciever     string
	Amount       int64
	Time_bill    string
	Task_id      int64
	Billing_type string
}
