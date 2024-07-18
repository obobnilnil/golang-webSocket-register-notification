package register

import (
	"database/sql"
	"webSocket_git/register/handlers"
	"webSocket_git/register/repositories"
	"webSocket_git/register/services"
	"webSocket_git/register/transactions"
	"webSocket_git/register/webSockets"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutesRegister(router *gin.Engine, db *sql.DB, conn *mongo.Database) {

	t := transactions.NewTransaction(conn)
	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServiceAdapter(r, t)
	h := handlers.NewHanerhandlerAdapter(s)

	router.POST("/api/registerChicCRM", h.RegisterChicCRMHandlers)

	router.GET("/ws/1", func(c *gin.Context) { // admin webSocket only
		webSockets.HandleConnections(c.Writer, c.Request)
	})
}
