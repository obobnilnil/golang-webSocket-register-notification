package handlers

import (
	"net/http"
	"webSocket_git/register/models"
	"webSocket_git/register/services"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	RegisterChicCRMHandlers(c *gin.Context)
}

type handlerAdapter struct {
	s services.ServicePort
}

func NewHanerhandlerAdapter(s services.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) RegisterChicCRMHandlers(c *gin.Context) {
	var loginData models.RegisterRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	companyID, err := h.s.RegisterChicCRMServices(loginData)
	if err != nil {
		switch err.Error() {
		case "email already exists", "mobile Phone already exists", "mobile phone must be 10 digits 089-XXX-XXXX", "username must be a valid email address":
			c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		case "Please fill in all the required information.":
			c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "companyID": companyID.CompanyID})
}
