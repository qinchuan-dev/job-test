package types

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

type DepositHistory struct {
	Id     string `json:"id"`
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
	Date   string `json:"date"`
	OpType string `json:"opType"`
	Memo   string `json:"memo"`
}

type SendHistory struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Denom    string `json:"denom"`
	Amount   string `json:"amount"`
	Date     string `json:"date"`
	Memo     string `json:"memo"`
}
