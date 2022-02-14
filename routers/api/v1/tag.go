package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"go-blog-demo/models"
	"go-blog-demo/pkg/err"
	"go-blog-demo/pkg/setting"
	"go-blog-demo/pkg/util"
)

func GetTags(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := c.Query("name"); name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		maps["state"] = com.StrTo(arg).MustInt()
	}

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": err.SUCCESS,
		"msg":  err.GetMsg(err.SUCCESS),
		"data": data,
	})
}

func AddTag(c *gin.Context) {

}

func EditTags(c *gin.Context) {

}

func DeleteTags(c *gin.Context) {

}
