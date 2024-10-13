package kategori

import (
	"go-learning-book/database/connection" // Ensure this imports the correct package
	"go-learning-book/middlewares"
	"go-learning-book/utils/common"

	"github.com/gin-gonic/gin"
)

// Initiator initializes all kategori routes
func Initiator(router *gin.Engine) {
	api := router.Group("/api/categories")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateKategoriRouter)       // Create
		api.GET("", GetAllKategoriRouter)        // Read (List)
		api.GET("/:id", GetKategoriRouter)       // Read (By ID)
		api.PUT("/:id", UpdateKategoriRouter)    // Update
		api.DELETE("/:id", DeleteKategoriRouter) // Delete
	}
}

// CreateKategoriRouter handles adding a new kategori
func CreateKategoriRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections) // Use the exported variable
	svc := NewKategoriService(repo)

	kategori, err := svc.CreateKategoriService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully added kategori data", kategori)
}

// GetAllKategoriRouter handles retrieving all kategori with pagination and search
func GetAllKategoriRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewKategoriService(repo)

	kategori, total, err := svc.GetAllKategoriService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	data := gin.H{"total": total, "data": kategori}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved all kategori data", data)
}

// GetKategoriRouter handles retrieving a kategori by ID
func GetKategoriRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewKategoriService(repo)

	kategori, err := svc.GetKategoriByIDService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully retrieved kategori data", kategori)
}

// UpdateKategoriRouter handles updating an existing kategori
func UpdateKategoriRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewKategoriService(repo)

	if err := svc.UpdateKategoriService(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully updated kategori data")
}

// DeleteKategoriRouter handles deleting a kategori by ID
func DeleteKategoriRouter(ctx *gin.Context) {
	repo := NewRepository(connection.DBConnections)
	svc := NewKategoriService(repo)

	if err := svc.DeleteKategoriService(ctx); err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully deleted kategori data")
}
