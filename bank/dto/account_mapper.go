package dto

import (
	"bank/repository"
)

func ToAccountResponse(a repository.Account) AccountResponse {
	return AccountResponse{
		AccountID:   a.AccountID,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}

func ToAccountResponses(accounts []repository.Account) []AccountResponse {
	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, ToAccountResponse(account))
	}

	return responses
}
