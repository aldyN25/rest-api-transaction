package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aldyN25/myproject/app/models"
	gormdb "github.com/aldyN25/myproject/pkg/gorm.db"
	"gorm.io/gorm"
)

func GetUsersByID(ctx context.Context, userID string) (*models.Users, error) {
	users := new(models.Users)

	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	err = db.Debug().Where("user_id = ?", userID).First(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUsersByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Users, error) {
	users := new(models.Users)

	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	err = db.Debug().Where("phone_number = ?", phoneNumber).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func Create(ctx context.Context, request models.Users) error {
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

func UpdateBalance(ctx context.Context, userID string, balance int) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}
	err = db.Debug().Table("users").Where("user_id = ?", userID).Updates(map[string]interface{}{
		"balance":    balance,
		"updated_at": time.Now(),
	}).Error

	if err != nil {
		return err
	}
	return nil
}
func UpdateUser(ctx context.Context, userID string, request models.Users) (*models.Users, error) {
	users := new(models.Users)
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Table("users").Where("user_id = ? ", userID).Updates(map[string]interface{}{
		"first_name": request.FirstName,
		"last_name":  request.LastName,
		"address":    request.Address,
		"updated_at": time.Now(),
	}).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}
