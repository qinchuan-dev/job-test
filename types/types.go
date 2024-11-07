package types

import "time"

type Customer struct {
	CustomerId string
}

type Deposit struct {
	CustomerId string
	Items      []DepositItem
}

type DepositItem struct {
	Denom  string
	Amount string
}

type OpType int

const (
	DEPOSIT OpType = iota
	WITHDRAW
)

type DepositHistoryItem struct {
	CustomerId string
	Denom      string
	Amount     string
	Date       time.Time
	Type       OpType
}

type SendHistoryItem struct {
	Serial   string
	Sender   string
	Receiver string
	Denom    string
	Amount   string
}
