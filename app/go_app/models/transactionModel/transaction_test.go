package transactionModel_test

import (
	"prj/models/transactionModel"
	"testing"
	"time"
)

const layout string = "2006/01/02"
const layoutQuery string = "2006-01-02"

func TestToTransaction(t *testing.T) {
	req := transactionModel.TransactionRequest{
		Description: "Test",
		Date:        "2023/10/25",
		Amount:      100.0,
	}

	trans, err := req.ToTransaction()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if trans.Description != req.Description {
		t.Errorf("Expected Description %v, got %v", req.Description, trans.Description)
	}

	expectedDate, _ := time.Parse(layout, req.Date)
	if trans.Date != expectedDate {
		t.Errorf("Expected Date %v, got %v", expectedDate, trans.Date)
	}

	if trans.Amount != req.Amount {
		t.Errorf("Expected Amount %v, got %v", req.Amount, trans.Amount)
	}
}

func TestGetEndTime(t *testing.T) {
	trans := transactionModel.Transaction{
		Date: time.Now(),
	}

	endTime := trans.GetEndTime()
	expectedEndTime := time.Now().Format(layoutQuery)

	if endTime != expectedEndTime {
		t.Errorf("Expected EndTime %v, got %v", expectedEndTime, endTime)
	}
}

func TestGetBegintime(t *testing.T) {
	trans := transactionModel.Transaction{
		Date: time.Now(),
	}

	beginTime := trans.GetBegintime()
	expectedBeginTime := time.Now().AddDate(0, -6, 0).Format(layoutQuery)

	if beginTime != expectedBeginTime {
		t.Errorf("Expected BeginTime %v, got %v", expectedBeginTime, beginTime)
	}
}

func TestToGlobalTransaction(t *testing.T) {
	trans := transactionModel.Transaction{
		Description: "Test",
		Date:        time.Now(),
		Amount:      100.0,
	}
	rate := float32(1.2)

	gbTrans := trans.ToGlobalTransaction(rate)

	if gbTrans.USDTransaction != trans {
		t.Errorf("Expected USDTransaction %v, got %v", trans, gbTrans.USDTransaction)
	}

	if gbTrans.ExchangeRate != rate {
		t.Errorf("Expected ExchangeRate %v, got %v", rate, gbTrans.ExchangeRate)
	}

	expectedConvertedAmount := trans.Amount * rate
	if gbTrans.ConvertedAmount != expectedConvertedAmount {
		t.Errorf("Expected ConvertedAmount %v, got %v", expectedConvertedAmount, gbTrans.ConvertedAmount)
	}
}
