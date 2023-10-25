package server

import (
	"fmt"
	"log"
	"net/http"
	transactionController "prj/controllers/transaction"
	"prj/repository"
	transactionrepository "prj/repository/transactionRepository"

	"github.com/gorilla/mux"
)

func StartServer(port string, repo *repository.Repository) {
	trsCtrl := transactionController.Controller{
		Repository: transactionrepository.NewTransactionRepo(repo),
	}
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/transactions/insert", trsCtrl.Store).Methods("POST")
	routes.HandleFunc("/transactions/{transaction_id}/country/{country_name}", trsCtrl.Retrieve).Methods("GET")
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}
