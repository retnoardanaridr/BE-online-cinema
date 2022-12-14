package routes

import (
	"BE-waysbuck-API/handlers"
	"BE-waysbuck-API/pkg/middleware"
	"BE-waysbuck-API/pkg/mysql"
	"BE-waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionID)).Methods("GET")
	r.HandleFunc("/transaction/user/{id}", middleware.Auth(h.GetTransUser)).Methods("GET")
	r.HandleFunc("/create-transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
