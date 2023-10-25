package transactionController

import (
	"encoding/json"
	"net/http"
	"prj/models/transactionModel"
	transactionrepository "prj/repository/transactionRepository"
	"prj/services/fiscaldata"

	"github.com/gorilla/mux"
)

var fdAPI fiscaldata.FiscalData = fiscaldata.NewFiscalData()

type Controller struct {
	Repository transactionrepository.TransactionRepoInterface
}

func (ctrl *Controller) Store(w http.ResponseWriter, r *http.Request) {
	var t transactionModel.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction, err := t.ToTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ctrl.Repository.InsertTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (ctrl *Controller) Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]
	countryName := vars["country_name"]

	t, err := ctrl.Repository.GetTransaction(transactionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rate, err := fdAPI.GetRate(countryName, t.GetBegintime(), t.GetEndTime())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gbTransaction := t.ToGlobalTransaction(rate)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gbTransaction)
}
