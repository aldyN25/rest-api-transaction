package transaction

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aldyN25/myproject/app/models"
	"github.com/aldyN25/myproject/app/repository"
	repo "github.com/aldyN25/myproject/app/repository"
	"github.com/aldyN25/myproject/pkg/utils"
	"github.com/gin-gonic/gin"
)

func TopUp(c *gin.Context) {
	req := models.TopUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	userIDString := fmt.Sprint(userID)

	dataUser, err := repository.GetUsersByID(c, userIDString)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	topUpID := utils.GenerateUUID()
	transactionID := utils.GenerateUUID()
	topUp := models.Transactions{
		TransactionID:   transactionID,
		TopUpID:         &topUpID,
		UserID:          dataUser.UserID,
		Amount:          req.Amount,
		TransactionType: "CREDIT",
		BalanceBefore:   dataUser.Balance,
		BalanceAfter:    dataUser.Balance + req.Amount,
		CreatedAt:       time.Now(),
	}

	err = repo.CreateTransaction(c, topUp)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	balance := dataUser.Balance + req.Amount
	err = repo.UpdateBalance(c, dataUser.UserID.String(), balance)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": topUp,
	})
}

func Payment(c *gin.Context) {
	req := models.PaymentRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	userIDString := fmt.Sprint(userID)

	dataUser, err := repository.GetUsersByID(c, userIDString)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	if dataUser.Balance < req.Amount {
		utils.RespondError(c, http.StatusBadRequest, "Balance is not enough")
		return
	}

	paymentID := utils.GenerateUUID()
	transactionID := utils.GenerateUUID()

	payment := models.Transactions{
		TransactionID:   transactionID,
		PaymentID:       &paymentID,
		UserID:          dataUser.UserID,
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		TransactionType: "DEBIT",
		BalanceBefore:   dataUser.Balance,
		BalanceAfter:    dataUser.Balance - req.Amount,
		CreatedAt:       time.Now(),
	}

	err = repo.CreateTransaction(c, payment)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	balance := dataUser.Balance - req.Amount
	err = repo.UpdateBalance(c, dataUser.UserID.String(), balance)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": payment,
	})
}

func Transfer(c *gin.Context) {
	req := models.TransferRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	userIDString := fmt.Sprint(userID)

	fromUser, err := repository.GetUsersByID(c, userIDString)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	targetUser, err := repository.GetUsersByID(c, req.TargetUser)
	if err != nil || targetUser == nil {
		msgErr := fmt.Sprintf("Failed to find target user: %v\n", req.TargetUser)
		utils.RespondError(c, http.StatusNotFound, msgErr)
		return
	}

	if fromUser.Balance < req.Amount {
		utils.RespondError(c, http.StatusBadRequest, "Balance is not enough")
		return
	}

	transferID := utils.GenerateUUID()
	transactionID := utils.GenerateUUID()
	transfer := models.Transactions{
		TransactionID:   transactionID,
		TransferID:      &transferID,
		UserID:          fromUser.UserID,
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		TransactionType: "DEBIT",
		BalanceBefore:   fromUser.Balance,
		BalanceAfter:    fromUser.Balance - req.Amount,
		CreatedAt:       time.Now(),
	}

	err = repo.CreateTransaction(c, transfer)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	balanceFromUser := fromUser.Balance - req.Amount
	balanceTargetUser := targetUser.Balance + req.Amount
	err = repo.UpdateBalance(c, fromUser.UserID.String(), balanceFromUser)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = repo.UpdateBalance(c, targetUser.UserID.String(), balanceTargetUser)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transfer,
	})
}

func GetAll(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDString := fmt.Sprint(userID)

	data, err := repository.GetAll(c, userIDString)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": data,
	})
}
