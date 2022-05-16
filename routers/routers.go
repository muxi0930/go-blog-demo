package routers

import (
	"go-blog-demo/pkg/setting"
	v1 "go-blog-demo/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.StaticFS("/files", http.Dir("."))
	r.Static("/web", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/staffs", v1.GetStaffs)
		apiv1.POST("/staff", v1.AddStaff)
		apiv1.PUT("/staffs/:id", v1.EditStaffs)
		apiv1.DELETE("/staffs/:id", v1.DeleteStaffs)
		apiv1.POST("/upload_header", v1.UploadHeader)
		apiv1.POST("/upload_certificate", v1.UploadCertificate)
	}
	return r
}
