package controller

import (
	"CodeHive/logger"
	"CodeHive/model"
	"CodeHive/service/code"
	"CodeHive/service/language"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CodeController struct {
	cs *code.CodeServiceImpl
	ls *language.LanguageServiceImpl
}

func NewCodeController(cs *code.CodeServiceImpl, ls *language.LanguageServiceImpl) *CodeController {
	return &CodeController{cs: cs, ls: ls}
}

func (cc *CodeController) CreateCode(c *gin.Context) {

	var cd model.Code

	if err := c.ShouldBindJSON(&c); err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	l, err := cc.ls.GetLanguageById(context.Background(), cd.LanguageID)
	if err != nil {
		logger.Log.Errorf("LanguageId not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	cd.Language = l

	_, err = cc.cs.Create(context.Background(), &cd)
	if err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":          "create code",
		"code information": cd,
	})
}
func (cc *CodeController) PutCode(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "update code",
	})
}
func (cc *CodeController) GetCodeById(c *gin.Context) {
	idStr := c.GetHeader("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cod, err := cc.cs.GetCodeById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"code": cod,
	})
}
func (cc *CodeController) GetListCode(c *gin.Context) {
	l, err := cc.cs.GetListCode(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code_list": l,
	})
}
