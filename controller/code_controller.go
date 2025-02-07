package controller

import (
	"CodeHive/logger"
	"CodeHive/model"
	"CodeHive/service/code"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeController struct {
	cs *code.CodeServiceImpl
}

func NewCodeController(sc *code.CodeServiceImpl) *CodeController {
	return &CodeController{sc}
}

func (cc *CodeController) CreateCode(c *gin.Context) {

	var code model.Code

	if err := c.ShouldBindJSON(&code); err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := cc.cs.Create(context.Background(), &code)
	if err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":          "create code",
		"code information": code,
	})
}
func (cc *CodeController) PutCode(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "update code",
	})
}
func (cc *CodeController) DeleteCode(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "delete successfull",
	})
}
func (cc *CodeController) GetCodes(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "delete successfull",
	})
}
