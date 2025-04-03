package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rudolfoborges/asapi/internal/entity"
	"github.com/rudolfoborges/asapi/internal/repository"
)

type CreateAccountInput struct {
	Name     string
	Email    string
	Password string
}

type AccountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s *AccountService) Create(ctx context.Context, input CreateAccountInput) error {
	existsByEmail, err := s.accountRepository.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return err
	}

	if existsByEmail {
		return fmt.Errorf("email %s already exists", input.Email)
	}

	account := &entity.Account{
		ID:        func() uuid.UUID { id, _ := uuid.NewV7(); return id }(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.accountRepository.Create(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
