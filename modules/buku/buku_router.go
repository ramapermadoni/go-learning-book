package buku

import (
	"go-learning-book/database/connection"
	"go-learning-book/middlewares"
	"go-learning-book/utils/common"

	"github.com/gin-gonic/gin"
)

// Initiator initializes all buku routes
func Initiator(router *gin.Engine) {
	api := router.Group("/api/books")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateBukuRouter)       // Create
		api.GET("", GetAllBukuRouter)        // Read (List)
		api.GET("/:id", GetBukuRouter)       // Read (By ID)
		api.PUT("/:id", UpdateBukuRouter)    // Update
		api.DELETE("/:id", DeleteBukuRouter) // Delete
	}
}

// CreateBukuRouter handles adding a new buku
func CreateBukuRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewBukuService(repo)

	buku, err := svc.CreateBukuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully added buku data", buku)
}

// GetAllBukuRouter handles retrieving all buku with pagination and search
func GetAllBukuRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewBukuService(repo)

	bukus, total, err := svc.GetAllBukuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	data := gin.H{"total": total, "data": bukus}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved all buku data", data)
}

// GetBukuRouter handles retrieving a buku by ID
func GetBukuRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewBukuService(repo)

	buku, err := svc.GetBukuByIDService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved buku data", buku)
}

// UpdateBukuRouter handles updating an existing buku
func UpdateBukuRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewBukuService(repo)

	if err := svc.UpdateBukuService(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully updated buku data")
}

// DeleteBukuRouter handles deleting a buku by ID
func DeleteBukuRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewBukuService(repo)

	if err := svc.DeleteBukuService(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully deleted buku data")
}
