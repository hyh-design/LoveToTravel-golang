package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ltt-gc/utils"
	"net/http"
)

/**
参考: https://blog.csdn.net/weixin_46272577/article/details/125175593
*/

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "RequestHeader-Authorization is empty!",
			})
			c.Abort()
			return
		}
		fmt.Println(token)
		mc, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的user信息保存到请求的上下文c上
		c.Set("name", mc.Name)
		c.Next() // 后续的处理函数可以用过c.Get("name")来获取当前请求的用户信息
	}
}

func HomeHandler(c *gin.Context) {
	name := c.MustGet("name").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": gin.H{"name": name},
	})
}
