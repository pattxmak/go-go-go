package handler

import (
	"bank/dto"
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{accService: accService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := dto.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode((&request))
	if err != nil {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}
	
	response, err := h.accService.NewAccount(customerID, request)
	if err != nil {
		handlerError(w, err)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {

	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	response, err := h.accService.GetAccounts(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}
