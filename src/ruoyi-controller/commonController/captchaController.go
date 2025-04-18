package commonController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	entity "ruoyi-gin/src/ruoyi-common/core/domain/entity/ajaxResult"
)

func GetCode(c *gin.Context) {
	ajaxResutlt := entity.NewAjaxResutlt()
	ajaxResutlt.SuccessDefault()
	//captchaEnabled := service.SelectCaptchaEnabled()
	ajaxResutlt.Result["captchaEnabled"] = false
	//不生成验证码
	c.SecureJSON(http.StatusOK, ajaxResutlt.Result)
}
