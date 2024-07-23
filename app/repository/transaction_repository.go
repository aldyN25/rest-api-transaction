package repository

import (
	"context"

	"github.com/aldyN25/myproject/app/models"
	gormdb "github.com/aldyN25/myproject/pkg/gorm.db"
)

func CreateTransaction(ctx context.Context, request models.Transactions) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Create(&request).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAll(ctx context.Context, userID string) ([]models.Transactions, error) {
	data := []models.Transactions{}
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	err = db.Where("user_id = ? ", userID).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
