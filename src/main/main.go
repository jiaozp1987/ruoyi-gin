package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/redis/go-redis/v9"
	_ "gopkg.in/ini.v1"
	"ruoyi-gin/src/ruoyi-common/core/middleware"
	"ruoyi-gin/src/ruoyi-common/core/routeConfig"
)

func main() {
	//fmt.Println("s")
	router := gin.Default()
	router.Use(middleware.CheckToken)
	//router.GET("/captchaImage", commonController.GetCode)
	//router.POST("/login", systemController.Login)
	//router.GET("/getInfo", systemController.GetInfo)
	router.GET(routeConfig.ROUTE_CONFIG["/captchaImage"][0].(string), routeConfig.ROUTE_CONFIG["/captchaImage"][1].(func(*gin.Context)))
	router.POST(routeConfig.ROUTE_CONFIG["/login"][0].(string), routeConfig.ROUTE_CONFIG["/login"][1].(func(*gin.Context)))
	router.GET(routeConfig.ROUTE_CONFIG["/getInfo"][0].(string), routeConfig.ROUTE_CONFIG["/getInfo"][1].(func(*gin.Context)))

	router.Run()
	//
	//hms := cryptor.HmacSha256("admin123", iniUtils.CRYPTOR_KEY)
	//fmt.Println(hms)
}

//import "github.com/duke-git/lancet/v2/cryptor"
//
//func main() {
//	// 加密（工作因子默认10）
//	hashedPwd, _ := cryptor.BcryptHash("rawPassword")
//
//	// 验证
//	isValid := cryptor.BcryptCheck("rawPassword", hashedPwd) // 返回bool
//}
