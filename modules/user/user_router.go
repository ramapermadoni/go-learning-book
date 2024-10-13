package user

import (
	"go-learning-book/database/connection" // Pastikan ini mengimpor paket yang benar
	"go-learning-book/middlewares"
	"go-learning-book/utils/common"

	"github.com/gin-gonic/gin"
)

// Initiator initializes all user routes
func Initiator(router *gin.Engine) {
	api := router.Group("/api/user")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateUserRouter)       // Create
		api.GET("", GetAllUserRouter)        // Read (List)
		api.GET("/:id", GetUserRouter)       // Read (By ID)
		api.PUT("/:id", UpdateUserRouter)    // Update
		api.DELETE("/:id", DeleteUserRouter) // Delete
	}
}

// CreateUserRouter handles adding a new user
func CreateUserRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections) // Gunakan variabel yang diekspor
	svc := NewUserService(repo)

	user, err := svc.CreateUser(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully added user data", user)
}

// GetAllUserRouter handles retrieving all users with pagination and search
func GetAllUserRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewUserService(repo)

	users, total, err := svc.GetAllUsers(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	data := gin.H{"total": total, "data": users}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved all user data", data)
}

// GetUserRouter handles retrieving a user by ID
func GetUserRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewUserService(repo)

	user, err := svc.GetUserByID(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved user data", user)
}

// UpdateUserRouter handles updating an existing user
func UpdateUserRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewUserService(repo)

	if err := svc.UpdateUser(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully updated user data")
}

// DeleteUserRouter handles deleting a user by ID
func DeleteUserRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewUserService(repo)

	if err := svc.DeleteUser(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully deleted user data")
}
