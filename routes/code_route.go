package routes

import (
	codeC "CodeHive/controller"
	codeR "CodeHive/repository/code"
	codeS "CodeHive/service/code"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCodeRoutes(r *gin.Engine, db *gorm.DB) {
	cr := codeR.NewCodeRepositoryImpl(db)
	cs := codeS.NewCodeServiceImpl(cr)
	cc := codeC.NewCodeController(cs)
	r.POST("/code/create", cc.CreateCode)
	r.GET("/code/list", cc.GetCodes)
}
