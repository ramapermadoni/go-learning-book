package main

import (
	"go-learning-book/config"
	"go-learning-book/database/connection"
	"go-learning-book/database/migration"
	"go-learning-book/modules/buku"
	"go-learning-book/modules/kategori"
	"go-learning-book/modules/user"
	"go-learning-book/utils/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Initiator()
	logger.Initiator()

	// Initialize the database connection
	connection.Initiator()

	// Pass the GORM connection to the migration function
	migration.Initiator(connection.DBConnections)

	// Initialize the router
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	user.UserInitiator(router)
	kategori.Initiator(router)
	buku.Initiator(router)
	user.Initiator(router)
	router.Run(":8080")
}
