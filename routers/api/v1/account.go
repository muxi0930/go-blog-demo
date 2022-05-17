package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"

	"go-blog-demo/models"
	"go-blog-demo/pkg/err"
	"go-blog-demo/pkg/util"
)

// type auth struct {
// 	Username string `valid:"Required; MaxSize(50)"`
// 	Password string `valid:"Required; MaxSize(50)"`
// }

func Login(c *gin.Context) {
	var auth = models.Auth{}
	c.ShouldBind(&auth)
	// valid := validation.Validation{}
	// a := auth{Username: username, Password: password}
	// ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := err.INVALID_PARAMS

	// if ok {
	if true {
		isExist := models.CheckAuth(auth.Username, auth.Password)
		if isExist {
			token, erro := util.GenerateToken(auth.Username, auth.Password)
			if erro != nil {
				code = err.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = err.SUCCESS
			}
		} else {
			code = err.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(err.SUCCESS),
		"data": data,
	})
}
