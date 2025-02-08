package controller

import (
	"CodeHive/logger"
	"CodeHive/model"
	"CodeHive/service/language"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LanguageController struct {
	ls *language.LanguageServiceImpl
}

func NewLanguageController(ls *language.LanguageServiceImpl) *LanguageController {
	return &LanguageController{ls: ls}
}
func (lc *LanguageController) CreateLanguage(c *gin.Context) {

	var l model.Language

	if err := c.ShouldBindJSON(&l); err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.ls.CreateLanguage(context.Background(), &l)
	if err != nil {
		logger.Log.Errorf("CreateLanguage err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":              "create language",
		"langauge information": l,
	})
}

func (lc *LanguageController) UpdateLanguage(c *gin.Context) {
	id := c.GetHeader("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is invalid").Error()})
	}

	var l model.Language

	if err := c.ShouldBindJSON(&l); err != nil {
		logger.Log.Errorf("CreateCode err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = lc.ls.UpdateLanguage(context.Background(), idInt, l)
	if err != nil {
		logger.Log.Errorf("UpdateLanguage err: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update language",
	})
}
