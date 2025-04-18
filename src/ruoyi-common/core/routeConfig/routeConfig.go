package routeConfig

import (
	"ruoyi-gin/src/ruoyi-controller/commonController"
	"ruoyi-gin/src/ruoyi-controller/systemController"
)

var (
	ROUTE_CONFIG map[string][3]interface{}
)

func init() {
	ROUTE_CONFIG = make(map[string][3]interface{})
	ROUTE_CONFIG["/captchaImage"] = [3]interface{}{"/captchaImage", commonController.GetCode, false}
	ROUTE_CONFIG["/login"] = [3]interface{}{"/login", systemController.Login, false}
	ROUTE_CONFIG["/getInfo"] = [3]interface{}{"/getInfo", systemController.GetInfo, true}
	ROUTE_CONFIG["/logout"] = [3]interface{}{"/logout", systemController.GetInfo, false}
	ROUTE_CONFIG["/getRouters"] = [3]interface{}{"/getRouters", systemController.GetRouters, true}
}
