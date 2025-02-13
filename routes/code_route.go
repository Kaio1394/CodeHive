package routes

import (
	codeC "CodeHive/controller"
	codeR "CodeHive/repository/code"
	languageR "CodeHive/repository/language"
	codeS "CodeHive/service/code"
	languageS "CodeHive/service/language"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCodeRoutes(r *gin.Engine, db *gorm.DB) {
	lr := languageR.NewLanguageRepositoryImpl(db)
	ls := languageS.NewLanguageServiceImpl(lr)

	cr := codeR.NewCodeRepositoryImpl(db)
	cs := codeS.NewCodeServiceImpl(cr)
	cc := codeC.NewCodeController(cs, ls)
	r.POST("/code/create", cc.CreateCode)
	r.GET("/code/list", cc.GetListCode)
	r.GET("/code/id", cc.GetCodeById)
}
