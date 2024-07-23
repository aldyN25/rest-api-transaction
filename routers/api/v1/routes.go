package v1

import (
	authController "github.com/aldyN25/myproject/app/controllers/auth"
	"github.com/aldyN25/myproject/app/controllers/transaction"
	usersController "github.com/aldyN25/myproject/app/controllers/users"
	"github.com/aldyN25/myproject/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	r := router.Group("/api/v1")
	auth := r.Group("/auth")

	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	transcation := r.Group("/transactions")
	{
		transcation.POST("/topup", utils.CheckAuth, transaction.TopUp)
		transcation.POST("/pay", utils.CheckAuth, transaction.Payment)
		transcation.POST("/transfer", utils.CheckAuth, transaction.Transfer)
		transcation.GET("/", utils.CheckAuth, transaction.GetAll)
	}

	users := r.Group("/users")
	// users.Use(utils.TokenVerify())
	{
		users.PUT("/profile", utils.CheckAuth, usersController.UpdateProfile)
	}
}
