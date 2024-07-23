package auth

import (
	"net/http"
	"time"

	"github.com/aldyN25/myproject/app/models"
	repo "github.com/aldyN25/myproject/app/repository"
	"github.com/aldyN25/myproject/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Register(c *gin.Context) {

	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			println("Field:", err.Field(), "Error:", err.Tag())
		}
	}

	// Cek Phone Number
	users, err := repo.GetUsersByPhoneNumber(c, user.PhoneNumber)
	if users != nil {
		utils.RespondError(c, http.StatusBadRequest, "Phone Number already registered")
		return
	}

	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.UserID = utils.GenerateUUID()
	hashedPIN, _ := bcrypt.GenerateFromPassword([]byte(user.PIN), bcrypt.DefaultCost)
	user.PIN = string(hashedPIN)
	user.Balance = 0
	user.CreatedAt = time.Now()

	err = repo.Create(c, user)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": user,
	})
}

func Login(c *gin.Context) {
	request := models.LoginRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			println("Field:", err.Field(), "Error:", err.Tag())
		}
	}

	// Cek Phone Number
	users, err := repo.GetUsersByPhoneNumber(c, request.PhoneNumber)
	if users == nil {
		utils.RespondError(c, http.StatusUnauthorized, "Phone Number and PIN doesn't match.")
		return
	}
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Phone Number and PIN doesn't match.")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.PIN), []byte(request.PIN)); err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Phone Number and PIN doesn't match.")
		return
	}

	accessToken, err := utils.GenerateAccessToken(users.UserID.String())
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(users.UserID.String())
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
