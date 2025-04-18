package middleware

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"ruoyi-gin/src/ruoyi-common/constants/constants"
	"ruoyi-gin/src/ruoyi-common/core/jwt"
	"ruoyi-gin/src/ruoyi-common/core/routeConfig"
)

func CheckToken(c *gin.Context) {

	url := c.Request.URL.String()
	fmt.Println(url)
	if routeConfig.ROUTE_CONFIG[url][2].(bool) {
		var token string
		// 找到token数据
		str := c.GetHeader("authorization")
		// 校验token
		token = strutil.ReplaceWithMap(str, map[string]string{constants.TOKEN_PREFIX: ""})
		loginuser, err := jwt.ParseToken(token)
		if err != nil {
			c.Abort()
			fmt.Println(err.Error())
		}
		c.Set("loginUser", loginuser)
	}

	c.Next()
}
