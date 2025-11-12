package service

import "bank/dto"

type AccountService interface {
	NewAccount(int, dto.NewAccountRequest) (*dto.AccountResponse, error)
	GetAccounts(int) ([]dto.AccountResponse, error)
}
