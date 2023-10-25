package transactionrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"prj/models/transactionModel"
	"prj/repository"
)

type TransactionRepo struct {
	Repository repository.RepositoryInterface
}

type TransactionRepoInterface interface {
	InsertTransaction(t transactionModel.Transaction) error
	GetTransaction(id string) (transactionModel.Transaction, error)
}

func (r TransactionRepo) InsertTransaction(t transactionModel.Transaction) error {
	sqlStatement := `
	INSERT INTO public.transaction (item_desc, purchase_date, amount, transaction_id)
	VALUES ($1, $2, $3, $4)`
	_, err := r.Repository.Exec(sqlStatement, t.Description, t.Date, t.Amount, t.Transaction_id)
	return err
}

func (r TransactionRepo) GetTransaction(id string) (transactionModel.Transaction, error) {
	row := r.Repository.QueryRow("SELECT transaction_id, item_desc, amount, purchase_date FROM public.transaction WHERE transaction_id = $1", id)

	var t transactionModel.Transaction
	err := row.Scan(&t.Transaction_id, &t.Description, &t.Amount, &t.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows were returned!")
			return transactionModel.Transaction{}, errors.New("transaction doesn't exist")
		} else {
			return transactionModel.Transaction{}, errors.New("something went wrong")
		}
	}
	return t, nil
}

func NewTransactionRepo(repo *repository.Repository) TransactionRepo {
	trsRepo := TransactionRepo{
		Repository: repo,
	}
	return trsRepo
}
