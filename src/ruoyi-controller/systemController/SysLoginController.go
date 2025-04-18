package systemController

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"ruoyi-gin/src/ruoyi-common/constants/constants"
	"ruoyi-gin/src/ruoyi-common/core/domain/entity/ajaxResult"
	lb "ruoyi-gin/src/ruoyi-domain/loginBody"
	lu "ruoyi-gin/src/ruoyi-domain/loginUser"
	service "ruoyi-gin/src/ruoyi-service/systemLoginService"
	"ruoyi-gin/src/ruoyi-service/systemMenuService"
)

func Login(c *gin.Context) {

	loginBody := lb.LoginBody{}
	if errA := c.ShouldBind(&loginBody); errA != nil {
		c.String(http.StatusOK, `the body should be loginBody`)
	}
	ajaxResutlt := ajaxResult.NewAjaxResutlt()
	//ajaxResutlt.SuccessDefault()
	token := service.Login(&loginBody)
	if strutil.IsBlank(token) {
		ajaxResutlt.ErrorDefault()
		c.SecureJSON(http.StatusOK, ajaxResutlt.Result)
	} else {
		ajaxResutlt.SuccessDefault()
		ajaxResutlt.Result[constants.TOKEN] = token
		fmt.Println(token)
		c.SecureJSON(http.StatusOK, ajaxResutlt.Result)
	}
}

func GetInfo(c *gin.Context) {
	usr, _ := c.Get("loginUser")
	loginUser := usr.(*lu.LoginUser)
	user, roles, permissions := service.GetInfo(loginUser)
	ajaxResutlt := ajaxResult.NewAjaxResutlt()
	ajaxResutlt.SuccessDefault()
	ajaxResutlt.Result["user"] = user
	ajaxResutlt.Result["roles"] = roles
	ajaxResutlt.Result["permissions"] = permissions
	c.SecureJSON(http.StatusOK, ajaxResutlt.Result)
}

func GetRouters(c *gin.Context) {
	usr, _ := c.Get("loginUser")
	loginUser := usr.(*lu.LoginUser)
	user, _, permissions := service.GetInfo(loginUser)
	sysuer := systemMenuService.SelectMenuTreeByUserId(&user, permissions)
	fmt.Println(sysuer)
}
