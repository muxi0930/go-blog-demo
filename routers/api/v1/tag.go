package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"go-blog-demo/models"
	"go-blog-demo/pkg/err"
	"go-blog-demo/pkg/setting"
	"go-blog-demo/pkg/util"
)

func GetStaffs(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := c.Query("name"); name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		maps["state"] = com.StrTo(arg).MustInt()
	}

	data["lists"] = models.GetStaffs(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetStaffTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": err.SUCCESS,
		"msg":  err.GetMsg(err.SUCCESS),
		"data": data,
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
