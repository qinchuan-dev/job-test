package pg

import (
	"job-test/types"
	"testing"
	"time"
)

func TestPostgres_GetDepositHistoryByCustomer(t *testing.T) {

}

func TestPostgres_GetSendHistoryByCustomer(t *testing.T) {

}

func TestPostgres_InsertDepositHistory(t *testing.T) {

}

func TestPostgres_InsertSendHistory(t *testing.T) {

	date := time.Now()

	item := types.DepositHistory{
		Id:     "id",
		Denom:  "denom",
		Amount: "1000",
		Date:   date.String(),
		OpType: string(rune(types.DEPOSIT)),
		Memo:   "memo",
	}
}
