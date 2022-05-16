package v1

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"go-blog-demo/models"
	"go-blog-demo/pkg/err"
	"go-blog-demo/pkg/util"
)

func GetStaffs(c *gin.Context) {
	maps := make(map[string]interface{})
	// data := make(map[string]interface{})
	if id := c.Query("id"); id != "" {
		maps["id"] = id
	}
	if name := c.Query("name"); name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		maps["state"] = com.StrTo(arg).MustInt()
	}

	lists := models.GetStaffs(util.GetPage(c), util.GetLimit(c), maps)
	total := models.GetStaffTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code":  err.LIST_SUCCESS,
		"msg":   err.GetMsg(err.SUCCESS),
		"data":  lists,
		"count": total,
	})
}

func AddStaff(c *gin.Context) {
	var staff = models.Staff{CreatedBy: "admin"}
	c.ShouldBind(&staff)
	models.AddStaff(staff)

	c.JSON(http.StatusOK, gin.H{
		"code": err.SUCCESS,
		"msg":  err.GetMsg(err.SUCCESS),
	})
}

func EditStaffs(c *gin.Context) {

}

func DeleteStaffs(c *gin.Context) {

}

// 接受上传图片接口
func UploadHeader(c *gin.Context) {
	file, erro := c.FormFile("file")
	fileName := file.Filename + "-" + com.ToStr(time.Now().Unix())
	if erro == nil {
		dst := path.Join("./static/upload/header", fileName)
		saveErr := c.SaveUploadedFile(file, dst)
		if saveErr == nil {
			log.Println("success save file to", dst)
			outDst := dst[7:]
			c.JSON(http.StatusOK, gin.H{
				"code": err.LIST_SUCCESS,
				"msg":  "",
				"data": outDst,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": err.ERROR,
		"msg":  "",
		"data": "",
	})
}

func UploadCertificate(c *gin.Context) {
	file, erro := c.FormFile("file")
	fileName := file.Filename + "-" + com.ToStr(time.Now().Unix())
	if erro == nil {
		dst := path.Join("./static/upload/certificate", fileName)
		saveErr := c.SaveUploadedFile(file, dst)
		if saveErr == nil {
			log.Println("success save file to", dst)
			outDst := dst[7:]
			c.JSON(http.StatusOK, gin.H{
				"code": err.LIST_SUCCESS,
				"msg":  "",
				"data": outDst,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": err.ERROR,
		"msg":  "",
		"data": "",
	})
}
