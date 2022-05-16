package util

import (
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"

    "go-blog-demo/pkg/setting"
)

func GetPage(c *gin.Context) int {
    result := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }
    return result
}

func GetLimit(c *gin.Context) int {
    limit, _ := com.StrTo(c.Query("limit")).Int()
    if limit < 10 {
        limit = 10
    }
    return limit
}
