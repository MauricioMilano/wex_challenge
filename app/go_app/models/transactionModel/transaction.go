package transactionModel

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type TransactionRequest struct {
	Description    string  `json:"description"`
	Date           string  `json:"date"`
	Amount         float32 `json:"amount"`
	Transaction_id string  `json:"id"`
}
type Transaction struct {
	Description    string    `json:"description"`
	Date           time.Time `json:"date"`
	Amount         float32   `json:"amount"`
	Transaction_id string    `json:"id"`
}
type GlobalTransaction struct {
	USDTransaction  Transaction
	ExchangeRate    float32
	ConvertedAmount float32
}

const layout string = "2006/01/02"
const layoutQuery string = "2006-01-02"

func (transactionReq TransactionRequest) ToTransaction() (Transaction, error) {
	transaction := Transaction{
		Transaction_id: uuid.New().String(),
		Amount:         float32(roundToDecimalPlaces(float64(transactionReq.Amount), 2)),
		Description:    transactionReq.Description,
	}
	date_parsed, err := time.Parse(layout, transactionReq.Date)
	if err != nil {
		return Transaction{}, err
	}
	transaction.Date = date_parsed
	return transaction, nil
}
func (transaction Transaction) GetEndTime() string {
	return transaction.Date.Format(layoutQuery)
}

func (transaction Transaction) GetBegintime() string {
	return transaction.Date.AddDate(0, -6, 0).Format(layoutQuery)
}
func (transaction Transaction) ToGlobalTransaction(rate float32) GlobalTransaction {
	gbTransaction := GlobalTransaction{
		USDTransaction:  transaction,
		ExchangeRate:    rate,
		ConvertedAmount: transaction.Amount * rate,
	}
	return gbTransaction
}

func roundToDecimalPlaces(num float64, decimalPlaces int) float64 {
	output := math.Round(num*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
	return output
}
