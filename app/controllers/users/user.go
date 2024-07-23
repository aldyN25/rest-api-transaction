package users

import (
	"fmt"
	"net/http"

	"github.com/aldyN25/myproject/app/models"
	"github.com/aldyN25/myproject/app/repository"
	repo "github.com/aldyN25/myproject/app/repository"
	"github.com/aldyN25/myproject/pkg/utils"
	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	req := models.UpdateUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.GetString("userID")
	userIDString := fmt.Sprint(userID)

	user := models.Users{}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address

	userUpdate, err := repo.UpdateUser(c, userIDString, user)
	if err != nil || userUpdate == nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	dataUser, err := repository.GetUsersByID(c, userIDString)
	if err != nil || dataUser == nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": dataUser,
	})
}
