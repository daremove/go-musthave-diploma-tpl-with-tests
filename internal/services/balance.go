package services

import (
	"context"
	"github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/database"
	"github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/models"
	"github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/utils"
	"sort"
)

type BalanceService struct {
	storage balanceStorage
}

type balanceStorage interface {
	FindAccrualFlow(ctx context.Context, userId string) (*[]database.AccrualFlowItemDB, error)

	CreateWithdrawal(ctx context.Context, orderId, userId string, amount float64) error

	FindWithdrawalFlow(ctx context.Context, userId string) (*[]database.WithdrawalFlowItemDB, error)
}

func NewBalanceService(storage balanceStorage) *BalanceService {
	return &BalanceService{storage: storage}
}

func (b *BalanceService) GetUserBalance(ctx context.Context, userId string) (models.Balance, error) {
	accrualFlow, err := b.storage.FindAccrualFlow(ctx, userId)

	if err != nil {
		return models.Balance{}, err
	}

	withdrawalFlow, err := b.storage.FindWithdrawalFlow(ctx, userId)

	if err != nil {
		return models.Balance{}, err
	}

	var current float64 = 0
	var withdrawn float64 = 0

	if accrualFlow != nil {
		for _, item := range *accrualFlow {
			current += item.Amount
		}
	}

	if withdrawalFlow != nil {
		for _, item := range *withdrawalFlow {
			withdrawn += item.Amount
		}
	}

	return models.Balance{Current: current - withdrawn, Withdrawn: withdrawn}, nil
}

func (b *BalanceService) CreateWithdrawal(ctx context.Context, orderId, userId string, amount float64) error {
	if err := b.storage.CreateWithdrawal(ctx, orderId, userId, amount); err != nil {
		return err
	}

	return nil
}

func (b *BalanceService) GetWithdrawalFlow(ctx context.Context, userId string) ([]models.WithdrawalFlowItem, error) {
	withdrawalFlow, err := b.storage.FindWithdrawalFlow(ctx, userId)

	if err != nil {
		return []models.WithdrawalFlowItem{}, err
	}

	if withdrawalFlow == nil {
		return []models.WithdrawalFlowItem{}, nil
	}

	var result []models.WithdrawalFlowItem

	for _, item := range *withdrawalFlow {
		result = append(result, models.WithdrawalFlowItem{
			OrderID:     item.OrderID,
			Sum:         item.Amount,
			ProcessedAt: utils.RFC3339Date{Time: item.ProcessedAt},
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ProcessedAt.Time.Before(result[j].ProcessedAt.Time)
	})

	return result, nil
}
