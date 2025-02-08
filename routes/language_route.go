package routes

import (
	"CodeHive/controller"
	lr "CodeHive/repository/language"
	ls "CodeHive/service/language"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLanguageRoutes(r *gin.Engine, db *gorm.DB) {
	cr := lr.NewLanguageRepositoryImpl(db)
	cs := ls.NewLanguageServiceImpl(cr)
	cc := controller.NewLanguageController(cs)
	r.POST("/language/create", cc.CreateLanguage)
	r.PUT("/language/update", cc.UpdateLanguage)
}
